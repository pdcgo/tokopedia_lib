/* eslint-disable @typescript-eslint/no-non-null-assertion */
/* eslint-disable react-hooks/exhaustive-deps */
import React, { Suspense, useEffect, useState } from "react"
import { Card, Checkbox, Divider, Pagination, Result, message } from "antd"
import { useRequest } from "../client"
import { useListProfileStore } from "../store/listProfile"
import { Flex, FlexColumn } from "../styled_components"
import { scroller } from "../utils/topScroller"

const UploadHeader = React.lazy(
    () => import("../component_sections/UploadHeader")
)
const ProfileCard = React.lazy(() => import("../components/ProfileCard"))

export default function Upload(props: {
    activePage?: string
}): React.ReactElement {
    const [
        profiles,
        markups,
        spins,
        collections,
        clipboard,
        pendingInit,
        error,
        totalData,
        initEffect,
        setClipboard,
        updateSingleProfile,
        updateAllProfileWith,
    ] = useListProfileStore((store) => [
        store.list,
        store.markups,
        store.spins,
        store.collections,
        store.clipboard,
        store.pendingInit,
        store.error,
        store.totalData,
        store.initEffect,
        store.setClipboard,
        store.updateSingleProfile,
        store.updateAllProfileWith,
        store.replaceAllProfile,
    ])

    const [query, setQuery] = useState({ page: 1, limit: 10, name: "" })
    const [useMapper, setUseMapper] = useState(false)
    const [showBottomPagination, setShowBottomPagination] = useState(false)
    const [messageApi, ctx] = message.useMessage()

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

    const { sender: getUseMapper } = useRequest("GetTokopediaMapperSetting", {
        onSuccess(data) {
            setUseMapper(data.use_mapper)
        },
    })
    const { sender: setUseMapperApi } = useRequest(
        "PutTokopediaMapperSetting",
        {
            onSuccess(data) {
                setUseMapper(data.use_mapper)
            },
        }
    )

    const { sender: deleterAccount } = useRequest("PostTokopediaAkunDelete", {
        onError(err) {
            message.error({ key: "error-delete", content: err.error })
        },
        onSuccess() {
            message.success({
                key: "success-delete",
                content: "Delete fulfilled!",
            })
            setQuery((q) => ({ ...q, name: "", page: 1 }))

            if (query.page == 1 && query.name == "") {
                initEffect(
                    query.limit,
                    (query.page - 1) * query.limit,
                    query.name
                )
            }
        },
    })

    useEffect(() => {
        if (pendingInit) {
            messageApi.loading({
                key: "load-accounts",
                content: "Loading accounts...",
            })
        } else {
            messageApi.destroy("load-accounts")
        }
    }, [pendingInit])

    useEffect(() => {
        if (props.activePage == "upload") {
            initEffect(query.limit, (query.page - 1) * query.limit, query.name)
            getUseMapper({ method: "get", path: "tokopedia/mapper/setting" })
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
        if (error !== null && !pendingInit) {
            return (
                <Flex style={{ justifyContent: "center", width: "100%" }}>
                    <Result status="error" title={error} subTitle={error} />
                </Flex>
            )
        } else if (!profiles.length && !pendingInit) {
            return (
                <Flex style={{ justifyContent: "center", width: "100%" }}>
                    <Result status="404" title="Data not found!" />
                </Flex>
            )
        }
    }

    function deleteSome() {
        const payload = profiles.filter((p) => p.isChecked).map((p) => p.id)
        deleterAccount({
            method: "post",
            path: "tokopedia/akun/delete",
            payload: {
                usernames: payload,
            },
        })
    }

    function updateAccount() {
        accountUpdater({
            method: "post",
            path: "tokopedia/akun/update",
            payload: {
                data: profiles.map((p) => ({
                    active_upload: p.isActive,
                    collection: p.colName,
                    count_upload: p.productCount,
                    hastag: "",
                    in_upload: p.isActive,
                    last_error: "",
                    lastup: 0,
                    limit_upload: p.limitUpload,
                    markup: p.markupName,
                    password: p.password,
                    secret: p.secret,
                    spin: p.spinName,
                    title_pattern: "",
                    username: p.emailOrUsername,
                })),
            },
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
                    disablePasteAll={clipboard === null}
                    checkedAll={
                        !pendingInit &&
                        profiles.length != 0 &&
                        profiles.every((p) => p.isChecked)
                    }
                    onChangeCheckedAll={(e) => {
                        if (e) {
                            updateAllProfileWith({ isChecked: true })
                        } else {
                            updateAllProfileWith({ isChecked: false })
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
                    onClickSetActive={() => {
                        updateAllProfileWith({ isActive: true })
                    }}
                    onClickSave={updateAccount}
                    loadingSave={pendingUpdateAccount}
                    loadingStartUpload={pendingUploadStarter}
                    onClickStartUpload={uploadAccount}
                    onClickPasteAll={() => {
                        if (clipboard) {
                            messageApi.success("Paste to All")
                            updateAllProfileWith({
                                markupName: clipboard.markupName,
                                spinName: clipboard.spinName,
                                colName: clipboard.colName,
                                limitUpload: clipboard.limitUpload,
                            })
                            setClipboard(null)
                        }
                    }}
                    indeterminate={
                        profiles.some((p) => p.isChecked) &&
                        !profiles.every((p) => p.isChecked)
                    }
                    disableRemoveAll={!profiles.some((p) => p.isChecked)}
                    onClickRemoveAll={deleteSome}
                />
            </Suspense>
            <Divider dashed style={{ margin: "5px 0" }} />
            {render()}
            <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "center",
                }}
                id="top-pagination"
            >
                {Boolean(profiles.length) && (
                    <>
                        <Pagination
                            pageSize={query.limit}
                            total={totalData}
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
                                    setQuery((q) => ({
                                        ...q,
                                        limit: size,
                                        page,
                                    }))
                                }
                            }}
                            showTotal={(tot) => `Total ${tot} profile`}
                        />
                        <Checkbox
                            checked={useMapper}
                            onChange={(e) => {
                                if (e.target.checked) {
                                    setUseMapperApi(
                                        {
                                            method: "put",
                                            path: "tokopedia/mapper/setting",
                                            payload: { use_mapper: true },
                                        },
                                        {
                                            onSuccess(data) {
                                                setUseMapper(data.use_mapper)
                                                message.info(
                                                    `Use category mapper ${
                                                        data.use_mapper
                                                            ? "ENABLED"
                                                            : "DISABLED"
                                                    }`
                                                )
                                            },
                                        }
                                    )
                                    return
                                }

                                setUseMapperApi(
                                    {
                                        method: "put",
                                        path: "tokopedia/mapper/setting",
                                        payload: { use_mapper: false },
                                    },
                                    {
                                        onSuccess(data) {
                                            setUseMapper(data.use_mapper)
                                            message.info(
                                                `Use category mapper ${
                                                    data.use_mapper
                                                        ? "ENABLED"
                                                        : "DISABLED"
                                                }`
                                            )
                                        },
                                    }
                                )
                            }}
                        >
                            Use Automatic Category Mapping
                        </Checkbox>
                    </>
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
                {profiles.map((profile, index) => (
                    <Suspense fallback={<Card loading />} key={profile.id}>
                        <ProfileCard
                            key={profile.id}
                            number={index + 1 + (query.page - 1) * query.limit}
                            clipboard={clipboard}
                            collections={collections}
                            markups={markups}
                            profile={profile}
                            spins={spins}
                            copyProfileFn={setClipboard}
                            updateSingleProfileFn={updateSingleProfile}
                            deleter={deleterAccount}
                        />
                    </Suspense>
                ))}
            </div>
            <div></div>
            {showBottomPagination && (
                <Pagination
                    pageSize={query.limit}
                    total={totalData}
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
                    showTotal={(tot) => `Total ${tot} profile`}
                />
            )}
        </FlexColumn>
    )
}
