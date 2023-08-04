/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import React, { Suspense, useEffect, useMemo, useState } from "react"
import { Card, Divider, Result } from "antd"
import { useRequest } from "../client"
import { Category } from "../client/sdk_types"
import { useListStore } from "../store/listMapper"
import { FlexColumn } from "../styled_components"

const TokopediaAccount = React.lazy(
    () => import("../components/TokopediaAccount")
)
const MapperHeader = React.lazy(
    () => import("../component_sections/MapperHeader")
)
const MapCard = React.lazy(() => import("../components/MapCard"))

export default function CategoryMapping(props: { activePage?: string }): React.ReactElement {
    const [showAsk, setShowAsk] = useState(false)
    const [list_, initEffect_, updateSingleList_, listPending_, reset_] =
        useListStore((state) => [
            state.list,
            state.initEffect,
            state.updateSingleList,
            state.pendingInitEffect,
            state.reset,
        ])

    const { sender, response: catListTokopedia } = useRequest(
        "GetTokopediaCategoryList"
    )
    const { sender: namespaceGetter, response } = useRequest(
        "GetV1ProductNamespaceAll"
    )

    const [selectedNamespacem, setSelectedNamespace] = useState<string | null>(
        null
    )

    useEffect(() => {
        sender(
            { method: "get", path: "tokopedia/category/list" },
            {
                onSuccess: (data) => {
                    if (!data) {
                        setShowAsk(true)
                    }
                },
            }
        )
        namespaceGetter(
            { method: "get", path: "v1/product/namespace_all" },
            {
                onSuccess(data) {
                    data?.forEach((nm) => {
                        if (nm.name !== "default") {
                            setSelectedNamespace(nm.name)
                        }
                    })
                },
            }
        )

        return () => reset_()
    }, [])

    useEffect(() => {
        if (selectedNamespacem && props.activePage == "category_map") {
            initEffect_(selectedNamespacem, catListTokopedia)
        }
    }, [selectedNamespacem, catListTokopedia])

    const namespaces = useMemo(() => {
        if (response) {
            return response.map((namespace) => ({
                label: namespace.name,
                value: namespace.name,
            }))
        }
        return []
    }, [response])

    const categories = () => {
        if (catListTokopedia) {
            const cb = (c: Category) => {
                const res = { label: c.name, value: c.id } as any

                if (c.children) {
                    res.children = c.children.map(cb)
                }

                return res
            }

            return (
                catListTokopedia.data.categoryAllListLite?.categories.map(cb) ||
                []
            )
        }

        return []
    }

    return (
        <FlexColumn>
            <Suspense fallback={<></>}>
                <TokopediaAccount
                    onFinish={() => setShowAsk(false)}
                    open={showAsk}
                    onCancel={() => setShowAsk(false)}
                />
            </Suspense>
            <Suspense fallback={<Card loading />}>
                <MapperHeader
                    list={list_}
                    namespace={selectedNamespacem}
                    namespaces={namespaces}
                    onChangeNamespace={setSelectedNamespace}
                    initEffect={initEffect_}
                    listCategoryTokopedia={catListTokopedia}
                />
            </Suspense>
            <Divider dashed style={{ marginBlock: "5px" }} />
            {!listPending_ && list_.length == 0 ? (
                <Result status="404" title="Data not found!" />
            ) : (
                <div
                    style={{
                        display: "grid",
                        gridTemplateColumns: "1fr ",
                        gap: "12px",
                    }}
                >
                    {list_.map((pc) => (
                        <Suspense
                            key={pc.shopeeCategoryId}
                            fallback={<Card loading />}
                        >
                            <MapCard
                                key={pc.shopeeCategoryId}
                                categoriesName={pc.shopeeCategoryName}
                                productCount={pc.productCount}
                                catsValue={pc.tokopediaCategoryIds}
                                onChangeCatsValue={(e) => {
                                    updateSingleList_(pc.shopeeCategoryId, {
                                        tokopediaCategoryIds: e,
                                    })
                                }}
                                optionsCats={categories()}
                            />
                        </Suspense>
                    ))}
                </div>
            )}
        </FlexColumn>
    )
}
