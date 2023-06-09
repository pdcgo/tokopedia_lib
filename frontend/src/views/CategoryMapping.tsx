/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-non-null-assertion */
import { Button, Card, Divider, Select, Typography } from "antd"
import { useEffect, useMemo, useState } from "react"
import { useRequest } from "../client"
import { Flex, FlexColumn } from "../styled_components"
import MapCard from "../components/MapCard"

export default function CategoryMapping(): React.ReactElement {
    const { sender } = useRequest("GetTokopediaCategoryList")
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
    const { sender: mapGetter } = useRequest("GetTokopediaMapperMap")

    const [selectedNamespacem, setSelectedNamespace] = useState<string | null>(
        null
    )

    useEffect(() => {
        sender({ method: "get", path: "tokopedia/category/list" })
        namespaceGetter({ method: "get", path: "v1/product/namespace_all" })
    }, [])

    useEffect(() => {
        if (selectedNamespacem) {
            mapGetter({
                method: "get",
                path: "tokopedia/mapper/map",
                params: { collection: selectedNamespacem },
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

    return (
        <FlexColumn>
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
            <Divider dashed style={{ marginBlock: "5px" }} />
            {/* <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "center",
                }}
            >
                <Flex>
                    <Typography.Text>
                        Total Category:{" "}
                        <Typography.Text style={{ fontWeight: 600 }}>
                            0
                        </Typography.Text>
                    </Typography.Text>
                    <Typography.Text>
                        Already Mapped:{" "}
                        <Typography.Text
                            style={{ fontWeight: 600, color: "green" }}
                        >
                            0
                        </Typography.Text>
                    </Typography.Text>
                </Flex>
            </Flex>
            <div></div> */}
            <div
                style={{
                    display: "grid",
                    gridTemplateColumns: "1fr 1fr",
                    gap: "7px",
                }}
            >
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
            </div>
        </FlexColumn>
    )
}
