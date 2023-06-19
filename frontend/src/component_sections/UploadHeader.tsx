import {
    CheckOutlined,
    DeleteOutlined,
    FilePptOutlined,
    ReloadOutlined,
    SaveOutlined,
    UploadOutlined,
} from "@ant-design/icons"
import { Button, Card, Checkbox, Input, message } from "antd"
import { Flex } from "../styled_components"
import { useRequest } from "../client"

export type UploadHeaderProps = {
    loadingSave?: boolean
    loadingStartUpload?: boolean
    disablePasteAll?: boolean
    disableRemoveAll?: boolean
    indeterminate?: boolean

    checkedAll?: boolean
    onChangeCheckedAll?: (v: boolean) => void

    nameQuery?: string
    onChangeNameQuery?: (name: string) => void

    onClickSetActive?: () => void
    onClickSave?: () => void
    onClickStartUpload?: () => void
    onClickPasteAll?: () => void
    onClickRemoveAll?: () => void
}

export default function UploadHeader(props: UploadHeaderProps) {
    const { sender: reset } = useRequest("PutTokopediaAkunResetAllCount", {
        onSuccess() {
            message.success({ key: "rss-scss", content: "Reset fulfilled" })
        },
        onError(err) {
            message.success({ key: "rss-err", content: err.error })
        },
    })

    return (
        <Card size="small" title="Setting Tokopedia Upload">
            <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "center",
                }}
            >
                <Checkbox
                    checked={props.checkedAll}
                    indeterminate={props.indeterminate}
                    onChange={(e) => {
                        props.onChangeCheckedAll?.(e.target.checked)
                    }}
                >
                    Select All
                </Checkbox>
                <Flex style={{ flex: 1 }}>
                    <Input
                        allowClear
                        placeholder="Search Profile..."
                        style={{ flex: 1 }}
                        value={props.nameQuery}
                        onChange={(e) =>
                            props.onChangeNameQuery?.(e.target.value)
                        }
                    />
                    <Button
                        onClick={props.onClickPasteAll}
                        icon={<FilePptOutlined rev="paste" />}
                        disabled={props.disablePasteAll}
                    >
                        Paste All
                    </Button>
                    <Button
                        disabled={props.disableRemoveAll}
                        icon={<DeleteOutlined rev="remove" />}
                        onClick={props.onClickRemoveAll}
                    >
                        Remove
                    </Button>
                    <Button
                        onClick={props.onClickSetActive}
                        icon={<CheckOutlined rev="active" />}
                    >
                        Active All
                    </Button>
                    <Button
                        onClick={() =>
                            reset({
                                method: "put",
                                path: "tokopedia/akun/reset_all_count",
                            })
                        }
                        icon={<ReloadOutlined rev="reset" />}
                    >
                        Reset All
                    </Button>
                    <Button
                        type="primary"
                        icon={<SaveOutlined rev="save" />}
                        onClick={props.onClickSave}
                        loading={props.loadingSave}
                        style={{
                            backgroundColor: "#C2418D",
                            boxShadow: "none",
                            color: "#fff",
                        }}
                    >
                        Save
                    </Button>

                    <Button
                        type="primary"
                        icon={<UploadOutlined rev="upload" />}
                        style={{ boxShadow: "none" }}
                        onClick={props.onClickStartUpload}
                        loading={props.loadingStartUpload}
                    >
                        Start Upload
                    </Button>
                </Flex>
            </Flex>
        </Card>
    )
}
