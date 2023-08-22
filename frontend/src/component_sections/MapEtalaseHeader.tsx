/* eslint-disable react-hooks/exhaustive-deps */
import React from "react"
import LabelInput from "../components/LabelInput"
import { Button, Card, Modal, Select, Typography } from "antd"
import { Flex, FlexColumn } from "../styled_components"
import { useRequest } from "../client"
import { DeleteFilled } from "@ant-design/icons"

export type Props = {
    collection?: string
    onChangeCollection?: (name: string) => void
    refetchFn?: () => void
}

export default function MapEtalaseHeader(props: Props) {
    const [openModal, setOpenModal] = React.useState(false)
    const { sender, response, pending } = useRequest("GetLegacyV1ProductNamespaceAll")
    const { sender: getEtalase, response: etalases } = useRequest(
        "GetTokopediaEtalaseMapListEtalase"
    )
    const { sender: deleter } = useRequest("DeleteTokopediaEtalaseMapDelete")

    const namespaces = React.useMemo(() => {
        if (!response) {
            return []
        }

        return response.map((r) => ({ value: r.name, label: r.name }))
    }, [response])

    function onDelete(name: string) {
        deleter(
            {
                method: "delete",
                path: "tokopedia/etalase_map/delete",
                params: { name },
            },
            {
                onSuccess() {
                    getEtalase({
                        method: "get",
                        path: "tokopedia/etalase_map/list_etalase",
                    })
                    props.refetchFn?.()
                },
            }
        )
    }

    React.useEffect(() => {
        if (openModal) {
            getEtalase({
                method: "get",
                path: "tokopedia/etalase_map/list_etalase",
            })
        }
    }, [openModal])

    React.useEffect(() => {
        const c = new AbortController()
        const signal = c.signal

        sender(
            { method: "get", path: "legacy/v1/product/namespace_all" },
            {
                signal,
                onSuccess(data) {
                    const f = data[0]
                    if (f) {
                        props.onChangeCollection?.(f.name)
                    }
                },
            }
        )

        return () => {
            c.abort()
            props.onChangeCollection?.("")
        }
    }, [])

    return (
        <Card
            size="small"
            title={
                <Typography.Text>Map Etalase From Collections</Typography.Text>
            }
        >
            <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "flex-end",
                }}
            >
                <LabelInput label="Collections :">
                    <Select
                        style={{ width: "300px" }}
                        placeholder="Choose collection"
                        options={namespaces}
                        loading={pending}
                        value={props.collection}
                        onChange={props.onChangeCollection}
                    ></Select>
                </LabelInput>
                <Button onClick={() => setOpenModal(true)}>
                    Delete Some Etalase
                </Button>
                <Modal
                    width={420}
                    footer={false}
                    closable={false}
                    open={openModal}
                    onCancel={() => setOpenModal(false)}
                    maskClosable
                    centered
                >
                    <Card
                        title="Etalase's Name List"
                        size="small"
                        type="inner"
                        extra={
                            <Button
                                size="small"
                                type="link"
                                onClick={() => setOpenModal(false)}
                            >
                                Back
                            </Button>
                        }
                    >
                        {etalases == null || !etalases.length ? (
                            <div css={{ textAlign: "center" }}>
                                <Typography.Text type="secondary">
                                    No list available
                                </Typography.Text>
                            </div>
                        ) : (
                            <div
                                style={{
                                    maxHeight: "300px",
                                    overflow: "auto",
                                }}
                            >
                                <FlexColumn style={{ paddingRight: 3 }}>
                                    {etalases
                                        .sort((a, v) =>
                                            a.etalase.toLowerCase() <
                                            v.etalase.toLowerCase()
                                                ? -1
                                                : 0
                                        )
                                        .map((e) => (
                                            <Card
                                                css={{
                                                    ":hover": {
                                                        cursor: "pointer",
                                                    },
                                                    "& button": {
                                                        transition:
                                                            "50ms ease !important",
                                                        opacity: 0,
                                                    },
                                                    ":hover button": {
                                                        opacity: 1,
                                                    },
                                                }}
                                                size="small"
                                                key={e.etalase}
                                            >
                                                <Flex
                                                    style={{
                                                        justifyContent:
                                                            "space-between",
                                                    }}
                                                >
                                                    <Typography.Text>
                                                        {e.etalase}
                                                    </Typography.Text>
                                                    <Button
                                                        ghost
                                                        size="small"
                                                        onClick={() =>
                                                            onDelete(e.etalase)
                                                        }
                                                        icon={
                                                            <DeleteFilled rev="knkcsd" />
                                                        }
                                                    />
                                                </Flex>
                                            </Card>
                                        ))}
                                </FlexColumn>
                            </div>
                        )}
                    </Card>
                </Modal>
            </Flex>
        </Card>
    )
}
