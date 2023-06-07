/* eslint-disable @typescript-eslint/no-non-null-assertion */
import { Flex, FlexColumn } from "../styled_components"
import { Button, Card, Divider, Pagination, Select, Typography } from "antd"
import MapCard from "../components/MapCard"
import { nanoid } from "nanoid"
import { useEffect, useState } from "react"
import { scroller } from "../utils/topScroller"

const topPaginationId = nanoid(7)

export default function CategoryMapping(): React.ReactElement {
    const [showBottomPagination, setShowBottomPagination] = useState(false)

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entry) => {
                if (!entry[0].isIntersecting) setShowBottomPagination(true)
                else setShowBottomPagination(false)
            },
            { threshold: [0] }
        )

        const el = document.getElementById(topPaginationId)

        if (el) {
            observer.observe(el)
        }

        return () => {
            if (el) {
                observer.unobserve(el)
            }
        }
    }, [])

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
                        />
                    </FlexColumn>
                    <Flex style={{ rowGap: "5px", justifyContent: "flex-end" }}>
                        <Button type="primary">Reset All</Button>
                        <Button type="primary">Save Mapping</Button>
                    </Flex>
                </Flex>
            </Card>
            <Divider dashed style={{ marginBlock: "5px" }} />
            <Flex id={topPaginationId} style={{ justifyContent: "flex-start" }}>
                <Pagination total={12} pageSize={2} showSizeChanger />
            </Flex>
            <div></div>
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
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
                <MapCard />
            </div>
            <div></div>
            {showBottomPagination && (
                <Flex style={{ justifyContent: "flex-start" }}>
                    <Pagination
                        onChange={() => {
                            scroller(true)
                        }}
                        total={12}
                        pageSize={2}
                        showSizeChanger
                    />
                </Flex>
            )}
        </FlexColumn>
    )
}
