import {
    DeleteOutlined,
    ReloadOutlined,
    UploadOutlined,
    CopyOutlined,
    FilePptOutlined,
} from "@ant-design/icons"
import {
    Card,
    Checkbox,
    Input,
    InputNumber,
    Select,
    Tooltip,
    Typography,
} from "antd"
import { Flex, FlexColumn } from "../styled_components"

export type ProfileCardProps = {
    number: number
    spins?: Array<{ data: string; name: string }>
    markups?: Array<string>
    collections?: Array<string>
    uploadCount?: number

    isActice?: boolean
    onChangeIsActive?: (v: boolean) => void

    markup?: string
    onChangeMarkup?: (v: string) => void

    spin?: string
    onChangeSpin?: (v: string) => void

    selected?: boolean
    onChangeSelected?: (v: boolean) => void

    limitUpload?: number
    onChangeLimitUpload?: (v: number | null) => void

    collection?: string
    onChangeCollection?: (v: string) => void

    username?: string
    onChangeUsername?: (username: string) => void

    password?: string
    onChangePassword?: (pass: string) => void

    onCopy?: () => void
    onPaste?: () => void
}

export default function ProfileCard(
    props: ProfileCardProps
): React.ReactElement {
    return (
        <Card
            title={
                <Checkbox
                    checked={props.selected}
                    onChange={(e) => props.onChangeSelected?.(e.target.checked)}
                    style={{ userSelect: "none" }}
                >
                    {props.number + ". "}
                    {props.username}
                </Checkbox>
            }
            hoverable
            size="small"
            type="inner"
            actions={[
                <Tooltip title="Upload" placement="bottom" showArrow={false}>
                    <UploadOutlined
                        style={{ color: "#FFA559" }}
                        rev={"upload"}
                        key="upload"
                    />
                </Tooltip>,
                <Tooltip title="Copy" placement="bottom" showArrow={false}>
                    <CopyOutlined
                        style={{ color: "#FFA559" }}
                        rev={"copy"}
                        key="copy"
                        onClick={props.onCopy}
                    />
                </Tooltip>,
                <Tooltip title="Paste" placement="bottom" showArrow={false}>
                    <FilePptOutlined
                        style={{ color: "#FFA559" }}
                        rev={"paste"}
                        key="paste"
                        onClick={props.onPaste}
                    />
                </Tooltip>,
                <Tooltip title="Reset" placement="bottom" showArrow={false}>
                    <ReloadOutlined
                        style={{ color: "#FFA559" }}
                        rev={"reset"}
                        key="reset"
                    />
                </Tooltip>,
                <Tooltip title="Remove" placement="bottom" showArrow={false}>
                    <DeleteOutlined
                        style={{ color: "#FFA559" }}
                        rev={"delete"}
                        key="delete"
                    />
                </Tooltip>,
            ]}
        >
            <FlexColumn>
                <Flex style={{ width: "100%" }}>
                    <FlexColumn style={{ flex: 1 }}>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Username :</Typography.Text>
                            <Input
                                value={props.username}
                                onChange={(e) =>
                                    props.onChangeUsername?.(e.target.value)
                                }
                                placeholder="username"
                            />
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Password :</Typography.Text>
                            <Input.Password
                                value={props.password}
                                onChange={(e) =>
                                    props.onChangePassword?.(e.target.value)
                                }
                                placeholder="⁎⁎⁎⁎⁎⁎⁎⁎"
                            />
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Upload Limit :</Typography.Text>
                            <InputNumber
                                value={props.limitUpload}
                                onChange={props.onChangeLimitUpload}
                                placeholder="1000"
                                style={{ width: "100%" }}
                            />
                        </FlexColumn>
                        <div></div>
                    </FlexColumn>
                    <FlexColumn style={{ flex: 1 }}>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Markup :</Typography.Text>
                            <Select
                                value={props.markup}
                                onChange={(v) => props.onChangeMarkup?.(v)}
                                placeholder="Choose Markup Data"
                            >
                                <Select.Option disabled value="">
                                    Markup Select
                                </Select.Option>
                                {props.markups?.map((markup) => (
                                    <Select.Option value={markup} key={markup}>
                                        {markup}
                                    </Select.Option>
                                ))}
                            </Select>
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Spin :</Typography.Text>
                            <Select
                                value={props.spin}
                                onChange={(v) => props.onChangeSpin?.(v)}
                                placeholder="Choose Spin Data"
                            >
                                <Select.Option disabled value="">
                                    Spin Select
                                </Select.Option>
                                {props.spins?.map((spin) => (
                                    <Select.Option
                                        value={spin.name}
                                        key={spin.data + spin.name}
                                    >
                                        {spin.name}
                                    </Select.Option>
                                ))}
                            </Select>
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text>Collection :</Typography.Text>
                            <Select
                                value={props.collection}
                                onChange={props.onChangeCollection}
                                placeholder="Choose Collection Data"
                            >
                                <Select.Option value="" disabled>
                                    Choose Collection
                                </Select.Option>
                                {props.collections?.map((collection) => (
                                    <Select.Option
                                        key={collection}
                                        value={collection}
                                    >
                                        {collection}
                                    </Select.Option>
                                ))}
                            </Select>
                        </FlexColumn>
                    </FlexColumn>
                </Flex>
                <Flex
                    style={{
                        justifyContent: "space-between",
                        width: "100%",
                    }}
                >
                    <Checkbox
                        checked={props.isActice}
                        onChange={(v) =>
                            props.onChangeIsActive?.(v.target.checked)
                        }
                        style={{ userSelect: "none" }}
                    >
                        Active
                    </Checkbox>
                    
                    <Typography.Text>Product Uploaded Count: {props.uploadCount || 0}</Typography.Text>
                </Flex>
            </FlexColumn>
        </Card>
    )
}
