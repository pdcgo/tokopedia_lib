/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-non-null-assertion */
import { Button, Card, Divider, Select, Typography } from "antd"
import { useEffect, useMemo, useState } from "react"
import { useRequest } from "../client"
import { Flex, FlexColumn } from "../styled_components"
import { Category } from "../client/sdk_types"
import MapCard from "../components/MapCard"
import TokopediaAccount from "../components/TokopediaAccount"

type List = {
    shopeeCats: string[]
    shopeeCatId: number
    productCount: number
    topedCatIds: (number|string)[]
}

export default function CategoryMapping(): React.ReactElement {
    const [showAsk, setShowAsk] = useState(false)
    const [list, setList] = useState<List[]>([])

    const { sender, response: catList } = useRequest(
        "GetTokopediaCategoryList",
        {
            onSuccess: (data) => {
                if (!data) {
                    setShowAsk(true)
                }
            },
        }
    )
    const { sender: namespaceGetter, response } = useRequest(
        "GetV1ProductNamespaceAll",
        {
            onSuccess(data) {
                data.forEach((nm) => {
                    if (nm.name !== "default") {
                        setSelectedNamespace(nm.name)
                    }
                })
            },
        }
    )
    const { sender: productCategoriesGetter } = useRequest(
        "GetV1ProductCategory",
        {
            onSuccess(data) {
                data.forEach((c) => {
                    setList((l) => [
                        ...l,
                        {
                            productCount: c.count,
                            shopeeCatId: c._id,
                            shopeeCats: c.name,
                            topedCatIds: [],
                        },
                    ])
                })
            },
        }
    )

    const [selectedNamespacem, setSelectedNamespace] = useState<string | null>(
        null
    )

    useEffect(() => {
        sender({ method: "get", path: "tokopedia/category/list" })
        namespaceGetter({ method: "get", path: "v1/product/namespace_all" })
    }, [])

    useEffect(() => {
        if (selectedNamespacem) {
            productCategoriesGetter({
                method: "get",
                path: "v1/product/category",
                params: {
                    is_public: false,
                    marketplace: "shopee",
                    namespace: selectedNamespacem,
                },
            })
        }
    }, [selectedNamespacem])

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
        if (catList) {
            const cb = (c: Category) => {
                const res = { label: c.name, value: c.id } as any

                if (c.children) {
                    res.children = c.children.map(cb)
                }

                return res
            }

            return catList.data.categoryAllListLite?.categories.map(cb) || []
        }

        return []
    }

    return (
        <FlexColumn>
            <TokopediaAccount
                onFinish={() => setShowAsk(false)}
                open={showAsk}
                onCancel={() => setShowAsk(false)}
            />
            <Card
                size="small"
                title={
                    <Typography.Text>
                        Map Category From Shopee to Tokopedia
                    </Typography.Text>
                }
            >
                <Flex
                    style={{
                        justifyContent: "space-between",
                        alignItems: "end",
                    }}
                >
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Collections :</Typography.Text>
                        <Select
                            style={{ width: "300px" }}
                            placeholder="Choose Collection"
                            value={selectedNamespacem}
                            onChange={setSelectedNamespace}
                            options={namespaces}
                        />
                    </FlexColumn>
                    <Flex style={{ rowGap: "5px", justifyContent: "flex-end" }}>
                        <Button>Use Suggest</Button>
                        <Button
                            style={{ backgroundColor: "#005246" }}
                            type="primary"
                        >
                            Reset All
                        </Button>
                        <Button type="primary">Save Mapping</Button>
                    </Flex>
                </Flex>
            </Card>
            <pre style={{maxHeight: 300}}>{JSON.stringify(list,null, 2)}</pre>
            <Divider dashed style={{ marginBlock: "5px" }} />
            <div
                style={{
                    display: "grid",
                    gridTemplateColumns: "1fr ",
                    gap: "7px",
                }}
            >
                {list.map((pc) => (
                    <MapCard
                        key={pc.shopeeCatId}
                        categoriesName={pc.shopeeCats}
                        productCount={pc.productCount}
                        catsValue={pc.topedCatIds}
                        onChangeCatsValue={e => {
                            setList(l => {
                                return l.map(ls => {
                                    if (ls.shopeeCatId == pc.shopeeCatId) {
                                        ls.topedCatIds = e || []
                                    }

                                    return ls
                                })
                            })
                        }}
                        optionsCats={categories()}
                    />
                ))}
            </div>
        </FlexColumn>
    )
}
