/* eslint-disable react-hooks/exhaustive-deps */
import {
    CheckOutlined,
    DeleteOutlined,
    FilePptOutlined,
    SaveOutlined,
    UploadOutlined,
} from "@ant-design/icons"
import {
    Button,
    Card,
    Checkbox,
    Divider,
    Input,
    Pagination,
    Result,
    message,
} from "antd"
import React, { useEffect, useState } from "react"
import { useRequest } from "../client"
import ProfileCard from "../components/ProfileCard"
import { Flex, FlexColumn } from "../styled_components"
import { scroller } from "../utils/topScroller"

export default function Upload(props: {
    activePage?: string
}): React.ReactElement {
    const [query, setQuery] = useState({ page: 1, limit: 10, name: "" })
    const [showBottomPagination, setShowBottomPagination] = useState(false)
    const [messageApi, ctx] = message.useMessage()

    const { sender, response, pending, error } = useRequest(
        "GetTokopediaAkunList"
    )

    useEffect(() => {
        if (pending) {
            messageApi.loading({
                key: "load-accounts",
                content: "Loading accounts...",
            })
        } else {
            messageApi.destroy("load-accounts")
        }
    }, [pending])

    useEffect(() => {
        if (props.activePage == "upload") {
            sender({
                method: "get",
                path: "/tokopedia/akun/list",
                params: {
                    limit: query.limit,
                    offset: (query.page - 1) * query.limit,
                    search: query.name,
                },
            })
        }
    }, [query.limit, query.name, query.page, props.activePage])

    useEffect(() => {
        const observer = new IntersectionObserver(
            function (entry) {
                if (!entry[0].isIntersecting) setShowBottomPagination(true)
                else setShowBottomPagination(false)
            },
            { threshold: [0] }
        )

        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        observer.observe(document.getElementById("top-pagination")!)

        return () => {
            // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
            observer.unobserve(document.getElementById("top-pagination")!)
        }
    }, [])

    function render() {
        if (error !== null && !pending) {
            return (
                <Flex style={{ justifyContent: "center", width: "100%" }}>
                    <Result
                        status="error"
                        title={error.msg}
                        subTitle={error.error}
                    />
                </Flex>
            )
        } else if (!response?.data.length && !pending) {
            return (
                <Flex style={{ justifyContent: "center", width: "100%" }}>
                    <Result status="404" title="Data not found!" />
                </Flex>
            )
        }
    }

    return (
        <FlexColumn>
            {ctx}
            <Card
                size="small"
                style={{ backgroundColor: "#fff6ea3e" }}
                title="Setting Tokopedia Upload"
            >
                <Flex
                    style={{
                        justifyContent: "space-between",
                        alignItems: "center",
                    }}
                >
                    <Checkbox>Select All</Checkbox>
                    <Flex style={{ flex: 1 }}>
                        <Input
                            allowClear
                            placeholder="Search Profile..."
                            style={{ flex: 1 }}
                            value={query.name}
                            onChange={(e) =>
                                setQuery((q) => ({
                                    ...q,
                                    page: 1,
                                    name: e.target.value,
                                }))
                            }
                        />
                        <Button icon={<FilePptOutlined rev="paste" />}>
                            Paste All
                        </Button>
                        <Button icon={<CheckOutlined rev="active" />}>
                            Set Active
                        </Button>
                        <Button danger icon={<DeleteOutlined rev="remove" />}>
                            Remove
                        </Button>
                        <Button
                            style={{
                                backgroundColor: "#FFA559",
                                boxShadow: "none",
                                color: "#454545",
                            }}
                            type="primary"
                            icon={<SaveOutlined rev="save" />}
                        >
                            Save
                        </Button>
                        <Button
                            type="primary"
                            icon={<UploadOutlined rev="upload" />}
                            style={{ boxShadow: "none" }}
                            onClick={() => {
                                messageApi.open({
                                    type: "loading",
                                    content: "Running process, please wait...",
                                    duration: 0,
                                    key: "loading",
                                })

                                setTimeout(() => {
                                    messageApi.destroy("loading")
                                }, 5000)
                            }}
                        >
                            Start Upload
                        </Button>
                    </Flex>
                </Flex>
            </Card>
            <Divider dashed style={{ margin: "5px 0" }} />
            {render()}
            <Flex style={{ justifyContent: "flex-start" }} id="top-pagination">
                {Boolean(response?.data.length) && (
                    <Pagination
                        pageSize={query.limit}
                        total={response?.pagination.count}
                        showSizeChanger
                        pageSizeOptions={[10, 20, 30, 40, 50, 75, 100]}
                        current={query.page}
                        onChange={(page, size) => {
                            if (query.limit !== size) {
                                setQuery((q) => ({
                                    ...q,
                                    limit: size,
                                    page: 1,
                                }))
                            } else {
                                setQuery((q) => ({ ...q, limit: size, page }))
                            }
                        }}
                    />
                )}
            </Flex>
            <div></div>
            <div
                style={{
                    display: "grid",
                    gap: "7px",
                    gridTemplateColumns: "1fr 1fr",
                }}
            >
                {response?.data.map((profile, index) => (
                    <ProfileCard
                        profile={profile}
                        key={profile.username}
                        number={index + 1 + (query.page - 1) * query.limit}
                    />
                ))}
            </div>
            <div></div>
            {showBottomPagination && (
                <Pagination
                    pageSize={query.limit}
                    total={response?.pagination.count}
                    showSizeChanger
                    pageSizeOptions={[10, 20, 30, 40, 50, 75, 100]}
                    current={query.page}
                    onChange={(page, size) => {
                        if (query.limit !== size) {
                            setQuery((q) => ({ ...q, limit: size, page: 1 }))
                        } else {
                            setQuery((q) => ({ ...q, limit: size, page }))
                        }

                        scroller()
                    }}
                />
            )}
        </FlexColumn>
    )
}