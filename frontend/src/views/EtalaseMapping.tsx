/* eslint-disable react-hooks/exhaustive-deps */
import React from "react"

import { a, useTrail } from "@react-spring/web"
import { Card, Divider, Result } from "antd"
import { useRequest } from "../client"
import { FlexColumn } from "../styled_components"

const EtalaseMapCard = React.lazy(() => import("../components/EtalaseMapCard"))
const Header = React.lazy(
    () => import("../component_sections/MapEtalaseHeader")
)
const loader = <Card loading size="small" />

export default function EtalaseMapping(props: { activePage?: string }) {
    const [namespace, setNamespace] = React.useState("")
    const {
        sender: listGetter,
        pending,
        response: listResponse,
    } = useRequest("GetTokopediaEtalaseMapList")
    const { sender: getEtalases, response: etalasesResponse } = useRequest(
        "GetTokopediaEtalaseMapListEtalase"
    )

    const list = React.useMemo(() => {
        if (listResponse) {
            return listResponse.data.sort((a, b) =>
                a.ShopeeCategoryName.join(" ") < b.ShopeeCategoryName.join(" ")
                    ? -1
                    : 0
            )
        }

        return []
    }, [listResponse])

    const etalases = React.useMemo(() => {
        if (etalasesResponse) {
            return etalasesResponse.map((e) => ({
                label: e.etalase,
                value: e.etalase,
            }))
        }

        return []
    }, [etalasesResponse])

    function refetch() {
        if (props.activePage == "etalase_map") {
            getEtalases({
                method: "get",
                path: "tokopedia/etalase_map/list_etalase",
            })

            if (namespace) {
                listGetter({
                    method: "get",
                    path: "tokopedia/etalase_map/list",
                    params: { namespace },
                })
            }
        }
    }

    React.useEffect(() => {
        if (props.activePage == "etalase_map") {
            getEtalases({
                method: "get",
                path: "tokopedia/etalase_map/list_etalase",
            })
        }
    }, [props.activePage])

    React.useEffect(() => {
        if (props.activePage == "etalase_map") {
            if (namespace)
                listGetter({
                    method: "get",
                    path: "tokopedia/etalase_map/list",
                    params: { namespace },
                })
        }
    }, [props.activePage, namespace])

    const trail = useTrail(list.length || 0, {
        config: { mass: 5, tension: 2000, friction: 200, duration: 70 },
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
                    refetchFn={refetch}
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
                        <a.div key={list[i].tokpedia_id} style={styles}>
                            <React.Suspense fallback={loader}>
                                <EtalaseMapCard
                                    item={list[i]}
                                    refetchFn={getEtalases}
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
