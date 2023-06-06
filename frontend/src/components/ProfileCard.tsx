import {
    DeleteOutlined,
    ReloadOutlined,
    UploadOutlined,
    CopyOutlined,
    FilePptOutlined,
} from "@ant-design/icons"
import { Card, Checkbox, Input, Select, Tooltip, Typography } from "antd"
import { Flex, FlexColumn } from "../styled_components"

export type Profile = {
    readonly username: string
    readonly password: string
}

export type ProfileCardProps = {
    profile: Profile
    number: number
    spins?: Array<{ data: string; name: string }>
    markups?: Array<string>

    isActice?: boolean
    onChangeIsActive?: (v: boolean) => void

    markup?: string
    onChangeMarkup?: (v: string) => void

    spin?: string
    onChangeSpin?: (v: string) => void

    selected?: boolean
    onChangeSelected?: (v: boolean) => void
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
                    {props.profile.username}
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
                <Tooltip title="Remove" placement="bottom" showArrow={false}>
                    <DeleteOutlined
                        style={{ color: "#FFA559" }}
                        rev={"delete"}
                        key="delete"
                    />
                </Tooltip>,
            ]}
        >
            <Flex style={{ width: "100%" }}>
                <FlexColumn style={{ flex: 1 }}>
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Username :</Typography.Text>
                        <Input
                            value={props.profile.username}
                            placeholder="username"
                        />
                    </FlexColumn>
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Password :</Typography.Text>
                        <Input.Password
                            value={props.profile.password}
                            placeholder="⁎⁎⁎⁎⁎⁎⁎⁎"
                        />
                    </FlexColumn>
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Upload Limit :</Typography.Text>
                        <Input placeholder="1000" />
                    </FlexColumn>
                    <div></div>
                    <Checkbox
                        checked={props.isActice}
                        onChange={(v) =>
                            props.onChangeIsActive?.(v.target.checked)
                        }
                        style={{ userSelect: "none" }}
                    >
                        Active
                    </Checkbox>
                </FlexColumn>
                <FlexColumn style={{ flex: 1 }}>
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Markup :</Typography.Text>
                        <Select
                            value={props.markup}
                            onChange={(v) => props.onChangeMarkup?.(v)}
                            placeholder="Choose Markup Data"
                        >
                            <option disabled value="">
                                Markup Select
                            </option>
                            {props.markups?.map((markup) => (
                                <option value={markup} key={markup}>
                                    {markup}
                                </option>
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
                            <option disabled value="">
                                Spin Select
                            </option>
                            {props.spins?.map((spin) => (
                                <option
                                    value={spin.name}
                                    key={spin.data + spin.name}
                                >
                                    {spin.name}
                                </option>
                            ))}
                        </Select>
                    </FlexColumn>
                    <FlexColumn style={{ rowGap: "5px" }}>
                        <Typography.Text>Collection :</Typography.Text>
                        <Select placeholder="Choose Collection Data" />
                    </FlexColumn>
                </FlexColumn>
            </Flex>
        </Card>
    )
}
