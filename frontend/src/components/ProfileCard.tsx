/* eslint-disable @typescript-eslint/no-explicit-any */
import {
    CopyOutlined,
    DeleteOutlined,
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
    message,
} from "antd"
import { SenderConfigs, UseQueryOptions } from "../client"
import { Response, SdkConfig } from "../client/sdk_types"
import {
    ListProfile,
    ListProfileActions,
    Selection,
} from "../store/listProfile"
import { Flex, FlexColumn } from "../styled_components"
import { Mode } from "../component_sections/UploadHeader"
import CollectionSelect from "./CollectionSelect"

export type ProfileCardProps = {
    updateSingleProfileFn: ListProfileActions["updateSingleProfile"]
    copyProfileFn: (profile: ListProfile) => void
    deleter: (
        config: SenderConfigs<
            SdkConfig["PostTokopediaAkunDelete"]["method"],
            SdkConfig["PostTokopediaAkunDelete"]["path"],
            SdkConfig["PostTokopediaAkunDelete"]["payload"],
            undefined
        >,
        senderOptions?: UseQueryOptions<Response, Response> | undefined
    ) => Promise<any>
    number: number
    profile: ListProfile
    clipboard: ListProfile | null
    markups: Selection[]
    spins: Selection[]
    collections: Selection[]
    mode: Mode
    manualCollections: Selection[]
}

export default function ProfileCard(
    props: ProfileCardProps
): React.ReactElement {
    return (
        <Card
            title={
                <Checkbox
                    checked={props.profile.isChecked}
                    onChange={(e) => {
                        props.updateSingleProfileFn(props.profile.id, {
                            isChecked: e.target.checked,
                        })
                    }}
                    style={{ userSelect: "none" }}
                >
                    {props.number + ". "}
                    {props.profile.id}
                </Checkbox>
            }
            hoverable
            size="small"
            type="inner"
            actions={[
                <Tooltip title="Copy" placement="bottom" showArrow={false}>
                    <CopyOutlined
                        style={{ color: "#FFA559" }}
                        rev={"copy"}
                        key="copy"
                        onClick={() => {
                            message.success(
                                `Copied profile: ${props.profile.id}`
                            )
                            props.copyProfileFn(props.profile)
                        }}
                    />
                </Tooltip>,
                <Tooltip title="Paste" placement="bottom" showArrow={false}>
                    <FilePptOutlined
                        style={{ color: "#FFA559" }}
                        rev={"paste"}
                        key="paste"
                        onClick={() => {
                            if (props.clipboard) {
                                message.info(
                                    `Paste from: ${props.clipboard.id}`
                                )
                                props.updateSingleProfileFn(props.profile.id, {
                                    limitUpload: props.clipboard.limitUpload,
                                    markupName: props.clipboard.markupName,
                                    spinName: props.clipboard.spinName,
                                    colName: props.clipboard.colName,
                                })
                            }
                        }}
                    />
                </Tooltip>,
                <Tooltip title="Remove" placement="bottom" showArrow={false}>
                    <DeleteOutlined
                        style={{ color: "#FFA559" }}
                        onClick={() =>
                            props.deleter({
                                method: "post",
                                path: "tokopedia/akun/delete",
                                payload: {
                                    usernames: [props.profile.id],
                                },
                            })
                        }
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
                            <Typography.Text type="secondary">
                                Username :
                            </Typography.Text>
                            <Input
                                value={props.profile.emailOrUsername}
                                onChange={(e) =>
                                    props.updateSingleProfileFn(
                                        props.profile.id,
                                        { emailOrUsername: e.target.value }
                                    )
                                }
                                placeholder="username"
                            />
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text type="secondary">
                                Password :
                            </Typography.Text>
                            <Input.Password
                                value={props.profile.password}
                                onChange={(e) =>
                                    props.updateSingleProfileFn(
                                        props.profile.id,
                                        { password: e.target.value }
                                    )
                                }
                                placeholder="⁎⁎⁎⁎⁎⁎⁎⁎"
                            />
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text type="secondary">
                                Upload Limit :
                            </Typography.Text>
                            <InputNumber
                                value={props.profile.limitUpload}
                                onChange={(e) =>
                                    props.updateSingleProfileFn(
                                        props.profile.id,
                                        { limitUpload: e || 0 }
                                    )
                                }
                                placeholder="1000"
                                style={{ width: "100%" }}
                            />
                        </FlexColumn>
                        <div></div>
                    </FlexColumn>
                    <FlexColumn style={{ flex: 1 }}>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text type="secondary">
                                Markup :
                            </Typography.Text>
                            <Select
                                value={props.profile.markupName}
                                onChange={(v) =>
                                    props.updateSingleProfileFn(
                                        props.profile.id,
                                        { markupName: v }
                                    )
                                }
                                placeholder="Choose Markup Data"
                            >
                                <Select.Option disabled value="">
                                    Markup Select
                                </Select.Option>
                                {props.markups?.map((markup) => (
                                    <Select.Option
                                        value={markup.value}
                                        key={markup.value}
                                    >
                                        {markup.label}
                                    </Select.Option>
                                ))}
                            </Select>
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text type="secondary">
                                Spin :
                            </Typography.Text>
                            <Select
                                value={props.profile.spinName}
                                onChange={(v) =>
                                    props.updateSingleProfileFn(
                                        props.profile.id,
                                        { spinName: v }
                                    )
                                }
                                placeholder="Choose Spin Data"
                            >
                                <Select.Option disabled value="">
                                    Spin Select
                                </Select.Option>
                                {props.spins?.map((spin) => (
                                    <Select.Option
                                        value={spin.value}
                                        key={spin.value}
                                    >
                                        {spin.label}
                                    </Select.Option>
                                ))}
                            </Select>
                        </FlexColumn>
                        <FlexColumn style={{ rowGap: "5px" }}>
                            <Typography.Text type="secondary">
                                Collection :
                            </Typography.Text>
                            <CollectionSelect
                                value={props.profile.colName}
                                mode={props.mode}
                                collections={props.collections}
                                manualCollections={props.manualCollections}
                                onChange={(colName) => props.updateSingleProfileFn(props.profile.id, { colName })}
                            />
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
                        checked={props.profile.isActive}
                        onChange={(v) =>
                            props.updateSingleProfileFn(props.profile.id, {
                                isActive: v.target.checked,
                            })
                        }
                        style={{ userSelect: "none" }}
                    >
                        Active
                    </Checkbox>

                    <Typography.Text>
                        Product Uploaded Count:{" "}
                        {props.profile.productCount || 0}
                    </Typography.Text>
                </Flex>
            </FlexColumn>
        </Card>
    )
}
