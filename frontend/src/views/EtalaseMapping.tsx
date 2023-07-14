/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Card, Popconfirm, Result, Tag, Typography } from "antd"
import React from "react"
import { useRequest } from "../client"
import { Flex, FlexColumn } from "../styled_components"
import { unwrapDeepCategory } from "../utils/unwrapDeepCategory"
import { DeleteOutlined, QuestionCircleOutlined } from "@ant-design/icons"

const AddEtalaseMapModal = React.lazy(
    () => import("../components/AddEtalaseMap")
)

const skeletons = <Card size="small" loading />

export default function EtalaseMapping() {
    const { sender, response, pending } = useRequest(
        "GetTokopediaEtalaseMapList"
    )
    const { sender: getCats, response: cats } = useRequest(
        "GetTokopediaCategoryList"
    )
    const { sender: delEtalase } = useRequest("DeleteTokopediaEtalaseMap")
    const [openModal, setOpenModal] = React.useState(false)

    const etalases = React.useMemo(() => {
        if (response) {
            if (cats?.data.categoryAllListLite) {
                const op = unwrapDeepCategory(cats)

                return response.data.map((et) => {
                    return {
                        name: et.etalase,
                        ids: et.cat_ids,
                        ids_names: et.cat_ids
                            .map((id) => {
                                for (const o of op) {
                                    if (o.value == id) {
                                        return o.label
                                    }
                                }
                            })
                            .filter(Boolean) as string[],
                    }
                })
            } else {
                return response.data.map((et) => {
                    return {
                        name: et.etalase,
                        ids: et.cat_ids,
                        ids_names: [] as string[],
                    }
                })
            }
        }

        return []
    }, [response])

    React.useEffect(() => {
        sender({ method: "get", path: "tokopedia/etalase_map/list" })
        getCats({ method: "get", path: "tokopedia/category/list" })
    }, [])

    return (
        <>
            <React.Suspense fallback={<></>}>
                <AddEtalaseMapModal
                    onFinish={() => {
                        setOpenModal(false)
                        sender({
                            method: "get",
                            path: "tokopedia/etalase_map/list",
                        })
                    }}
                    open={openModal}
                    onCancel={() => setOpenModal(false)}
                />
            </React.Suspense>
            {pending ? (
                skeletons
            ) : etalases.length ? (
                <FlexColumn style={{ alignItems: "start" }}>
                    <Button onClick={() => setOpenModal(true)} type="primary">
                        Add Map
                    </Button>
                    <div
                        style={{
                            display: "grid",
                            gridTemplateColumns:
                                "repeat(3, minmax(100px, 1fr))",
                            gap: "8px",
                            alignItems: "start",
                            width: "100%",
                        }}
                    >
                        {etalases.map((et) => (
                            <Card
                                hoverable
                                style={{ width: "100%", position: "relative" }}
                                size="small"
                                key={et.name}
                            >
                                <Popconfirm
                                    title="Delete the task"
                                    description="Are you sure to delete this task?"
                                    icon={
                                        <QuestionCircleOutlined
                                            style={{ color: "red" }}
                                            rev="frdel"
                                        />
                                    }
                                    onConfirm={() => {
                                        delEtalase({
                                            method: "delete",
                                            path: "tokopedia/etalase_map",
                                            params: { name: et.name },
                                        }, {
                                            onSuccess: () => {
                                                sender({
                                                    method: "get",
                                                    path: "tokopedia/etalase_map/list",
                                                })
                                            }
                                        })
                                    }}
                                >
                                    <Button
                                        size="small"
                                        type="ghost"
                                        icon={<DeleteOutlined rev="del" />}
                                        style={{
                                            position: "absolute",
                                            zIndex: 99,
                                            right: 8,
                                            top: 6,
                                        }}
                                    />
                                </Popconfirm>
                                <FlexColumn style={{ rowGap: 0 }}>
                                    <Typography.Title level={5}>
                                        {et.name}
                                    </Typography.Title>
                                    <Flex
                                        style={{
                                            flexWrap: "nowrap",
                                            width: "100%",
                                            overflow: "hidden",
                                            gap: "6px 0px",
                                        }}
                                    >
                                        {et.ids_names
                                            .slice(0, 2)
                                            .map((name) => {
                                                return (
                                                    <Tag
                                                        style={{
                                                            fontSize: "11.5px",
                                                        }}
                                                        key={name}
                                                        color="cyan"
                                                    >
                                                        {name}
                                                    </Tag>
                                                )
                                            })}
                                        {et.ids.length > 2 && (
                                            <Tag
                                                style={{
                                                    fontSize: "11.5px",
                                                }}
                                            >
                                                +{et.ids.length - 2} Other
                                            </Tag>
                                        )}
                                    </Flex>
                                </FlexColumn>
                            </Card>
                        ))}
                    </div>
                </FlexColumn>
            ) : (
                <Result
                    status="404"
                    title="Etalase mapping not available"
                    extra={
                        <Button
                            onClick={() => setOpenModal(true)}
                            type="primary"
                        >
                            Add Map
                        </Button>
                    }
                />
            )}
        </>
    )
}
