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
}

export default function CheckBotAsk(props: CheckBotAskProps & ModalProps) {
    const [filename, setFilename] = useState("")
    return (
        <Modal
            width={390}
            footer={false}
            closable={false}
            centered
            onCancel={() => {
                if (filename) {
                    props.onFinish(filename)
                } else {
                    props.onFinish("cekbot.csv")
                }
            }}
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
                        Default filename: cekbot.csv
                    </Typography.Text>
                    <Divider dashed style={{ marginBlock: 4 }} />
                    <Button
                        onClick={() => {
                            if (filename) {
                                props.onFinish(filename)
                            } else {
                                props.onFinish("cekbot.csv")
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
