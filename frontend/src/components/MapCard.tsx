import {
    CopyOutlined,
    FilePptOutlined,
    ReloadOutlined,
} from "@ant-design/icons"
import { Breadcrumb, Card, Cascader, Tooltip, Typography } from "antd"
import { Flex, FlexColumn } from "../styled_components"

export default function MapCard(): React.ReactElement {
    return (
        <Card
            extra={
                <span>
                    <strong>120</strong> Product
                </span>
            }
            size="small"
            type="inner"
            hoverable
            title={
                <Breadcrumb
                    separator=">"
                    items={[
                        {
                            title: "Fashion Pria",
                        },
                        {
                            title: "Atasan",
                        },
                        {
                            title: "Kemeja Flanel",
                        },
                    ]}
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
                    <Cascader showSearch style={{ width: "100%" }} />
                </Flex>
            </FlexColumn>
        </Card>
    )
}
