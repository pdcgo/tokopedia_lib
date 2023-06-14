/* eslint-disable @typescript-eslint/no-non-null-assertion */
/* eslint-disable react-hooks/exhaustive-deps */

import { Card, Divider, Pagination, Result, message } from "antd"
import React, { Suspense, useEffect, useState } from "react"
import { useRequest } from "../client"
import { AkunItem } from "../client/sdk_types"
import { Flex, FlexColumn } from "../styled_components"
import { scroller } from "../utils/topScroller"

const UploadHeader = React.lazy(
    () => import("../component_sections/UploadHeader")
)
const ProfileCard = React.lazy(() => import("../components/ProfileCard"))

export default function Upload(props: {
    activePage?: string
}): React.ReactElement {
    const [query, setQuery] = useState({ page: 1, limit: 10, name: "" })
    const [showBottomPagination, setShowBottomPagination] = useState(false)
    const [payload, setPayload] = useState<AkunItem[]>([])
    const [selectedAccounts, setSelectedAccounts] = useState<Array<string>>([])

    const [accountClip, setAccountClip] = useState<AkunItem | null>(null)

    const [messageApi, ctx] = message.useMessage()

    const { sender, response, pending, error } = useRequest(
        "GetTokopediaAkunList",
        {
            onSuccess(data) {
                setPayload([...data.data])
            },
        }
    )
    const { sender: spinGetter, response: spinData } =
        useRequest("GetApiSettingSpin")
    const { sender: markupGetter, response: markupData } =
        useRequest("GetApiListMarkup")

    const { sender: accountUpdater, pending: pendingUpdateAccount } =
        useRequest("PostTokopediaAkunUpdate", {
            onSuccess: () => message.success("Account list updated :)"),
            onError: (e) => message.error(JSON.stringify(e)),
        })

    const { sender: uploadStarter, pending: pendingUploadStarter } = useRequest(
        "GetTokopediaUploadStart",
        {
            onSuccess: () => message.success("Account list upload start :)"),
            onError: (e) => message.error(JSON.stringify(e)),
        }
    )

    const { sender: collectionGetter, response: collections } = useRequest(
        "GetV1ProductNamespaceAll"
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
            setSelectedAccounts([])
            sender({
                method: "get",
                path: "tokopedia/akun/list",
                params: {
                    limit: query.limit,
                    offset: (query.page - 1) * query.limit,
                    search: query.name,
                },
            })
            spinGetter({ method: "get", path: "api/settingSpin" })
            markupGetter({ method: "get", path: "api/listMarkup" })
            collectionGetter({
                method: "get",
                path: "v1/product/namespace_all",
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

        observer.observe(document.getElementById("top-pagination")!)

        return () => {
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

    function updateAccount() {
        accountUpdater({
            method: "post",
            path: "tokopedia/akun/update",
            payload: { data: payload },
        })
    }

    function uploadAccount() {
        uploadStarter({
            method: "get",
            path: "tokopedia/upload/start",
        })
    }

    return (
        <FlexColumn>
            {ctx}
            <Suspense fallback={<Card loading />}>
                <UploadHeader
                    checkedAll={payload.length === selectedAccounts.length}
                    onChangeCheckedAll={(e) => {
                        if (e) {
                            setSelectedAccounts([
                                ...payload.map((p) => p.username),
                            ])
                        } else {
                            setSelectedAccounts([])
                        }
                    }}
                    nameQuery={query.name}
                    onChangeNameQuery={(e) =>
                        setQuery((q) => ({
                            ...q,
                            page: 1,
                            name: e,
                        }))
                    }
                    onClickSetActive={() =>
                        setPayload((p) =>
                            p.map((payload) => {
                                if (
                                    selectedAccounts.includes(payload.username)
                                ) {
                                    payload.active_upload = true
                                }

                                return payload
                            })
                        )
                    }
                    onClickSave={updateAccount}
                    loadingSave={pendingUpdateAccount}
                    loadingStartUpload={pendingUploadStarter}
                    onClickStartUpload={uploadAccount}
                    onClickPasteAll={() => {
                        if (accountClip) {
                            messageApi.success("Paste to All")
                            setPayload((p) =>
                                p.map((payload) => ({
                                    ...payload,
                                    limit_upload: accountClip.limit_upload,
                                    markup: accountClip.markup,
                                    spin: accountClip.spin,
                                    collection: accountClip.collection,
                                }))
                            )
                        }
                    }}
                />
            </Suspense>
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
                {payload.map((profile, index) => (
                    <Suspense
                        fallback={<Card loading />}
                        key={profile.username}
                    >
                        <ProfileCard
                            key={profile.username}
                            number={index + 1 + (query.page - 1) * query.limit}
                            spins={spinData?.titlePool}
                            markups={markupData?.data}
                            collections={collections?.map(
                                (collection) => collection.name
                            )}
                            isActice={profile.active_upload}
                            uploadCount={profile.count_upload}
                            onChangeIsActive={(ck) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.active_upload = ck
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            markup={profile.markup}
                            onChangeMarkup={(mk) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.markup = mk
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            spin={profile.spin}
                            onChangeSpin={(sp) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.spin = sp
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            selected={selectedAccounts.includes(profile.username)}
                            onChangeSelected={(sl) => {
                                if (sl) {
                                    setSelectedAccounts((acc) => [
                                        ...acc,
                                        profile.username,
                                    ])
                                } else {
                                    setSelectedAccounts((acc) =>
                                        acc.filter((ac) => ac !== profile.username)
                                    )
                                }
                            }}
                            limitUpload={profile.limit_upload}
                            onChangeLimitUpload={(lm) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.limit_upload = lm || 0
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            collection={profile.collection}
                            onChangeCollection={(cl) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.collection = cl
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            username={profile.username}
                            onChangeUsername={(username_) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.username = username_
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            password={profile.password}
                            onChangePassword={(password_) => {
                                setPayload((p) =>
                                    p.map((payload) => {
                                        if (payload.username == profile.username) {
                                            payload.password = password_
                                        }
    
                                        return payload
                                    })
                                )
                            }}
                            onCopy={() => {
                                setAccountClip(profile)
                                messageApi.destroy("copiedprofile")
                                messageApi.success({
                                    content: "Profile Copied",
                                    key: "copiedprofile",
                                    duration: 1.5,
                                })
                            }}
                            onPaste={() => {
                                if (accountClip) {
                                    setPayload((py) =>
                                        py.map((payload) => {
                                            if (
                                                payload.username == profile.username
                                            ) {
                                                payload.collection =
                                                    accountClip.collection
                                                payload.markup = accountClip.markup
                                                payload.spin = accountClip.spin
                                                payload.limit_upload =
                                                    accountClip.limit_upload
                                            }
    
                                            return payload
                                        })
                                    )
                                }
                            }}
                        />
                    </Suspense>
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
