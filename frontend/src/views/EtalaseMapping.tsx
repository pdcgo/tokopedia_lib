/* eslint-disable react-hooks/exhaustive-deps */
import React from "react"

import { Card, Divider, Result } from "antd"
import { useRequest } from "../client"
import { useListStore } from "../store/listMapper"
import { FlexColumn } from "../styled_components"

const EtalaseMapCard = React.lazy(() => import("../components/EtalaseMapCard"))
const Header = React.lazy(
    () => import("../component_sections/MapEtalaseHeader")
)
const loader = <Card loading size="small" />

export default function EtalaseMapping(props: { activePage?: string }) {
    const [namespace, setNamespace] = React.useState("")
    const [initEffect, list, pending] = useListStore((s) => [
        s.initEffect,
        s.list,
        s.pendingInitEffect,
    ])
    const { sender, response } = useRequest("GetTokopediaCategoryList")

    React.useEffect(() => {
        const controller = new AbortController()

        sender(
            { method: "get", path: "tokopedia/category/list" },
            { signal: controller.signal }
        )

        return () => {
            controller.abort()
        }
    }, [])

    React.useEffect(() => {
        const controller = new AbortController()

        if (namespace != "" && response && props.activePage == "etalase_map") {
            initEffect(namespace, response, controller.signal)
        }

        return () => {
            controller.abort()
        }
    }, [namespace, response])

    return (
        <FlexColumn>
            <React.Suspense fallback={loader}>
                <Header
                    collection={namespace}
                    onChangeCollection={setNamespace}
                />
            </React.Suspense>
            <Divider dashed style={{ marginBlock: "5px" }} />
            <div
                style={{
                    display: "grid",
                    gap: "10px",
                    gridTemplateColumns: "repeat(2, 1fr)",
                }}
            >
                {!pending &&
                    namespace != "" &&
                    list.map((l) => (
                        <React.Suspense
                            key={l.shopeeCategoryId}
                            fallback={loader}
                        >
                            <EtalaseMapCard
                                productCount={l.productCount}
                                shopeeCatNames={l.shopeeCategoryName}
                                key={l.shopeeCategoryId}
                            />
                        </React.Suspense>
                    ))}
            </div>
            {!pending && namespace != "" && list.length < 1 && (
                <Result status="404" title="Data not found!" />
            )}
        </FlexColumn>
    )
}
