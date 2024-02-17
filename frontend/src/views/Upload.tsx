/* eslint-disable @typescript-eslint/no-non-null-assertion */
/* eslint-disable react-hooks/exhaustive-deps */
import { Card, Checkbox, Divider, Pagination, Result, message } from "antd"
import React, { Suspense, useEffect, useState } from "react"
import { useRequest } from "../client"
import { useQuery } from "../client/newapisdk"
import { useMutation } from "../client/sdk_mutation"
import { ManualQuery } from "../component_sections/UploadHeader"
import { useListProfileStore, Selection } from "../store/listProfile"
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
    const [upquery, setUpquery] = useState<ManualQuery>({
        mode: "shopee",
        reset: false,
        one_to_multi: false,
        limit: 0,
    })
    const [collections, setCollection] = useState<Selection[]>([])
    const [manualCollections, setManualCollection] = useState<Selection[]>([])
    const [useMapper, setUseMapper] = useState(false)
    const [showBottomPagination, setShowBottomPagination] = useState(false)
    const [messageApi, ctx] = message.useMessage()

    const { sender: accountUpdater, pending: pendingUpdateAccount } =
        useRequest("PostTokopediaAkunUpdate", {
            onSuccess: () => message.success("Account list updated :)"),
            onError: (e) => message.error(JSON.stringify(e)),
        })

    const { send: uploadShopee, pending: pendingUploadShopee } = useQuery("GetTokopediaUploadShopee", {})
    const { send: uploadTokped, pending: pendingUploadTokped } = useQuery("GetTokopediaUploadTokopedia", {})
    const { send: uploadTokpedManual, pending: pendingUploadTokpedManual } = useQuery("GetUploadV6ManualToTokopedia", {})
    const { send: uploadJakmall, pending: pendingUploadJakmall } = useQuery("GetUploadV6JakmallToTokopedia", {})
    const { mutate: setUseMapperApi } = useMutation("PutTokopediaMapperSetting", {})
    const { send: getCollectionList } = useQuery("GetLegacyV1ProductNamespaceAll")
    const { send: getManualCollectionList } = useQuery("GetPdcsourceCollectionList")

    const { sender: getUseMapper } = useRequest("GetTokopediaMapperSetting", {
        onSuccess(data) {
            setUseMapper(data.use_mapper)
        },
    })

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
        getCollectionList({
            query: {
                is_public: false,
                kota: "",
                namespace: "",
                pmax: 0,
                pmin: 0,
                marketplace: upquery.mode,
                use_empty_city: false,
            },
            onSuccess(res) {
                setCollection(res.map((val) => ({
                    label: val.name,
                    value: val.name
                })))
            },
        })
    }, [upquery.mode])

    useEffect(() => {

        getManualCollectionList({
            query: {
                page: 1,
                limit: 999999,
            },
            onSuccess(data) {
                const cols = data.data.reduce<Selection[]>((res, val) => {
                    val && res.push({
                        label: val.name,
                        value: val.name
                    })
                    return res
                }, [])
                setManualCollection(cols)
            },
        })

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
        switch (upquery.mode) {
            case "tokopedia":
                uploadTokped({
                    onSuccess: () => message.success("Account list upload tokopedia :)"),
                    onError: (e) => message.error(JSON.stringify(e)),
                })
                break

            case "tokopedia_manual":
                uploadTokpedManual({
                    query: {
                        base: "./",
                        use_mapper: useMapper,
                        reset: upquery.reset,
                        one_to_multi: upquery.one_to_multi,
                        limit: upquery.limit,
                    },
                    onSuccess: () => message.success("Account list upload tokopedia manual :)"),
                    onError: (e) => message.error(JSON.stringify(e)),
                })
                break

            case "jakmall":
                uploadJakmall({
                    query: {
                        base: "./",
                        use_mapper: useMapper,
                    },
                    onSuccess: () => message.success("Account list upload jakmall :)"),
                    onError: (e) => message.error(JSON.stringify(e)),
                })
                break

            default:
                uploadShopee({
                    onSuccess: () => message.success("Account list upload shopee :)"),
                    onError: (e) => message.error(JSON.stringify(e)),
                })
        }
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
                    loadingStartUpload={
                        pendingUploadShopee ||
                        pendingUploadTokped ||
                        pendingUploadTokpedManual ||
                        pendingUploadJakmall
                    }
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
                    upquery={upquery}
                    onUploadQueryChange={setUpquery}
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
                                setUseMapperApi({
                                    onSuccess(data) {
                                        setUseMapper(data.use_mapper)
                                        message.info(
                                            `Use category mapper ${data.use_mapper
                                                ? "ENABLED"
                                                : "DISABLED"
                                            }`
                                        )
                                    },
                                }, { use_mapper: e.target.checked })
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
                            mode={upquery.mode}
                            number={index + 1 + (query.page - 1) * query.limit}
                            clipboard={clipboard}
                            collections={collections}
                            manualCollections={manualCollections}
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
