import {
    CheckOutlined,
    DeleteOutlined,
    FilePptOutlined,
    SaveOutlined,
    UploadOutlined,
} from "@ant-design/icons"
import { Button, Card, Checkbox, Input } from "antd"
import { Flex } from "../styled_components"

export type UploadHeaderProps = {
    loadingSave?: boolean
    loadingStartUpload?: boolean
    disablePasteAll?: boolean

    checkedAll?: boolean
    onChangeCheckedAll?: (v: boolean) => void

    nameQuery?: string
    onChangeNameQuery?: (name: string) => void

    onClickSetActive?: () => void
    onClickSave?: () => void
    onClickStartUpload?: () => void
    onClickPasteAll?: () => void
}

export default function UploadHeader(props: UploadHeaderProps) {
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
                        onClick={props.onClickSetActive}
                        icon={<CheckOutlined rev="active" />}
                    >
                        Active All
                    </Button>
                    <Button disabled icon={<DeleteOutlined rev="remove" />}>
                        Remove
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
