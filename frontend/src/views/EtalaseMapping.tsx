/* eslint-disable react-hooks/exhaustive-deps */
import React from "react"

import { a, useTrail } from "@react-spring/web"
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
    const { sender: getEtalases, response: etalasesResponse } = useRequest(
        "GetTokopediaEtalaseMapListEtalase"
    )

    const etalases = React.useMemo(() => {
        if (etalasesResponse) {
            return etalasesResponse.map((e) => ({
                label: e.etalase,
                value: e.etalase,
            }))
        }

        return []
    }, [etalasesResponse])

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
        if (props.activePage == "etalase_map")
            getEtalases({
                method: "get",
                path: "tokopedia/etalase_map/list_etalase",
            })
    }, [props.activePage])

    React.useEffect(() => {
        const controller = new AbortController()

        if (namespace != "" && response && props.activePage == "etalase_map") {
            initEffect(namespace, response, controller.signal)
        }

        return () => {
            controller.abort()
        }
    }, [namespace, response])

    const trail = useTrail(list.length, {
        config: { mass: 5, tension: 2000, friction: 200, duration: 80 },
        opacity: list.length ? 1 : 0,
        y: list.length ? 0 : -10,
        from: { opacity: 0, y: -10 },
    })

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
                    trail.map(({ ...styles }, i) => (
                        <a.div key={list[i].shopeeCategoryId} style={styles}>
                            <React.Suspense fallback={loader}>
                                <EtalaseMapCard
                                    productCount={list[i].productCount}
                                    shopeeCatNames={list[i].shopeeCategoryName}
                                    etalases={etalases}
                                />
                            </React.Suspense>
                        </a.div>
                    ))}
            </div>
            {!pending && namespace != "" && list.length < 1 && (
                <Result status="404" title="Data not found!" />
            )}
        </FlexColumn>
    )
}
