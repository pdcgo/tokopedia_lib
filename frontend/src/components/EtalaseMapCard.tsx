/* eslint-disable react-hooks/exhaustive-deps */
import { Breadcrumb, Button, Card, Divider, Input, Select } from "antd"
import { Flex, FlexColumn } from "../styled_components"

export type Props = {
    shopeeCatNames?: Array<string>
    productCount?: number
    style?: React.CSSProperties
    etalases?: {value: string, label: string}[]
}

export default function EtalaseMapCard(props: Props) {
    return (
        <Card size="small" type="inner" style={props.style}>
            <FlexColumn
                style={{ justifyContent: "space-between", height: "100%" }}
            >
                <Flex style={{ justifyContent: "space-between", flex: 1 }}>
                    <Breadcrumb
                        items={props.shopeeCatNames?.map((f) => ({ title: f }))}
                        separator=">"
                    />
                    <div style={{ width: "20px" }}></div>
                    <span style={{ flexShrink: 0 }}>
                        <strong>{props.productCount || 0}</strong> Product
                    </span>
                </Flex>
                <Select
                    placeholder="Choose etalase"
                    dropdownRender={(menu) => {
                        return (
                            <>
                                {menu}
                                <Divider style={{ marginBlock: 4 }}></Divider>
                                <Flex style={{ columnGap: 5 }}>
                                    <Input placeholder="Create new etalase..." />
                                    <Button type="primary">Add</Button>
                                </Flex>
                            </>
                        )
                    }}
                    options={props.etalases}
                />
            </FlexColumn>
        </Card>
    )
}
