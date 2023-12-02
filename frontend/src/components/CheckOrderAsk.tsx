import {
    Button,
    Card,
    Divider,
    Input,
    Modal,
    ModalProps,
    Typography,
} from "antd"
import { FlexColumn } from "../styled_components"
import { useState } from "react"

export type CheckBotAskProps = {
    onFinish: (name: string) => void
    onCancel: () => void
}

export default function CheckOrderAsk(props: CheckBotAskProps & ModalProps) {
    const [filename, setFilename] = useState("")
    return (
        <Modal
            width={390}
            footer={false}
            closable={false}
            centered
            {...props}
        >
            <Card title="File Target Name" size="small" type="inner">
                <FlexColumn style={{ rowGap: 5 }}>
                    <Input
                        value={filename}
                        onChange={(e) => setFilename(e.target.value)}
                        addonAfter=".csv"
                        placeholder="Boleh dikosongi"
                    />
                    <Typography.Text>
                        Default filename: cekorder_tokopedia.csv
                    </Typography.Text>
                    <Divider dashed style={{ marginBlock: 4 }} />
                    <Button
                        onClick={() => {
                            if (filename) {
                                props.onFinish(filename + ".csv")
                            } else {
                                props.onFinish("cekorder_tokopedia.csv")
                            }
                        }}
                        type="primary"
                    >
                        Run
                    </Button>
                </FlexColumn>
            </Card>
        </Modal>
    )
}
