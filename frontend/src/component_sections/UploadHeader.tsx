import {
    CheckOutlined,
    DeleteOutlined,
    FilePptOutlined,
    ReloadOutlined,
    SaveOutlined,
    UploadOutlined,
} from "@ant-design/icons"
import { Button, Card, Checkbox, Input, InputNumber, Select, Space, message } from "antd"

import { useRequest } from "../client"
import { Flex } from "../styled_components"

export type Mode = "shopee" | "tokopedia" | "tokopedia_manual" | "jakmall"

export interface ManualQuery {
    mode: Mode
    reset: boolean
    one_to_multi: boolean
    limit: number
}

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

    upquery: ManualQuery
    onUploadQueryChange?: (query: ManualQuery) => void

    onClickSetActive?: () => void
    onClickSave?: () => void
    onClickStartUpload?: () => void
    onClickPasteAll?: () => void
    onClickRemoveAll?: () => void
}

export default function UploadHeader(props: UploadHeaderProps) {
    const { upquery, onUploadQueryChange } = props;
    const isManual = upquery.mode === "tokopedia_manual"

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
                <Input
                    allowClear
                    placeholder="Search Profile..."
                    style={{ flex: 1 }}
                    value={props.nameQuery}
                    onChange={(e) =>
                        props.onChangeNameQuery?.(e.target.value)
                    }
                />
                <Flex style={{ flex: 1 }}>

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
                </Flex>
            </Flex>
            <Flex
                style={{
                    alignItems: "center",
                    marginTop: 10
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
                <Flex style={{ flex: 1, justifyContent: "end" }}>
                    <Space>
                        <Checkbox
                            disabled={!isManual}
                            style={{ fontWeight: 300 }}
                            onChange={(e) => onUploadQueryChange?.({ ...upquery, reset: e.target.checked })}
                        >Reset Mapper</Checkbox>
                        <Checkbox
                            disabled={!isManual}
                            style={{ fontWeight: 300 }}
                            onChange={(e) => onUploadQueryChange?.({ ...upquery, one_to_multi: e.target.checked })}
                        >One to Multi</Checkbox>
                        <span>Limit :</span>
                        <InputNumber
                            value={upquery.limit}
                            disabled={!isManual}
                            style={{ width: 150 }}
                            onChange={(v) => onUploadQueryChange?.({ ...upquery, limit: v || 1 })}
                        />
                        <span>Mode :</span>
                        <Select
                            value={upquery.mode}
                            style={{ minWidth: 200 }}
                            options={[
                                { value: "shopee", label: "Shopee" },
                                { value: "tokopedia", label: "Tokopedia" },
                                { value: "tokopedia_manual", label: "Tokopedia Manual" },
                                { value: "jakmall", label: "Jakmall" },
                            ]}
                            onChange={(mode) => onUploadQueryChange?.({ ...upquery, mode })}
                        />
                    </Space>
                    <Button
                        type="primary"
                        icon={<UploadOutlined rev="upload" />}
                        style={{ boxShadow: "none" }}
                        onClick={() => props.onClickStartUpload?.()}
                        loading={props.loadingStartUpload}
                    >
                        Start Upload
                    </Button>
                </Flex>
            </Flex>
        </Card >
    )
}
