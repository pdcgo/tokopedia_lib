/* eslint-disable @typescript-eslint/no-non-null-assertion */
import {
    CopyOutlined,
    FilePptOutlined,
    ReloadOutlined,
} from "@ant-design/icons"
import {
    Breadcrumb,
    Card,
    Cascader,
    CascaderProps,
    Tooltip,
    Typography,
} from "antd"
import { Flex, FlexColumn } from "../styled_components"

export type MapCardProps = {
    categoriesName?: string[]
    productCount?: number
    optionsCats?: CascaderProps["options"]
    catsValue?: (number|string)[]
    onChangeCatsValue?: (j: (number|string)[]) => void
}

export default function MapCard(props: MapCardProps): React.ReactElement {
    return (
        <Card
            extra={
                <span>
                    <strong>{props.productCount}</strong> Product
                </span>
            }
            size="small"
            type="inner"
            hoverable
            title={
                <Breadcrumb
                    separator="/"
                    items={props.categoriesName?.map((cat) => ({ title: cat }))}
                />
            }
            actions={[
                <Tooltip title="Copy" placement="bottom" showArrow={false}>
                    <CopyOutlined
                        style={{ color: "#FFA559" }}
                        rev={"copy"}
                        key="copy"
                    />
                </Tooltip>,
                <Tooltip title="Paste" placement="bottom" showArrow={false}>
                    <FilePptOutlined
                        style={{ color: "#FFA559" }}
                        rev={"paste"}
                        key="paste"
                    />
                </Tooltip>,
                <Tooltip title="Reset" placement="bottom" showArrow={false}>
                    <ReloadOutlined
                        style={{ color: "#FFA559" }}
                        rev={"reset"}
                        key="reset"
                    />
                </Tooltip>,
            ]}
        >
            <FlexColumn style={{ rowGap: "5px" }}>
                <Typography.Text>Map to :</Typography.Text>
                <Flex style={{ columnGap: "7px" }}>
                    <Cascader
                        value={props.catsValue}
                        onChange={props.onChangeCatsValue}
                        options={props.optionsCats}
                        showSearch
                        style={{ width: "100%" }}
                    />
                </Flex>
            </FlexColumn>
        </Card>
    )
}
