import {
    Button,
    Card,
    Divider,
    Input,
    Modal,
    ModalProps,
    Switch,
    Typography,
} from "antd"
import { FlexColumn } from "../styled_components"
import { useState } from "react"
import CheckOrderDateRange from "./CheckOrderDateRange"
import CheckOrderSelectStatus from "./CheckOrderSelectStatus"
import { CheckOrderConfig } from "../client/newapisdk"

export type CheckBotAskProps = {
    onFinish: (name: string, config: CheckOrderConfig) => void
    onCancel: () => void
}

export default function CheckOrderAsk(props: CheckBotAskProps & ModalProps) {
    const [name, setName] = useState("")
    const [config, setConfig] = useState<CheckOrderConfig>({
        useDateRange: false,
        startDate: "",
        endDate: "",
        useStatus: false,
        statusKeys: [],
    })

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
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        addonAfter=".csv"
                        placeholder="Boleh dikosongi"
                    />
                    <Typography.Text>
                        Default filename: cekorder_tokopedia.csv
                    </Typography.Text>
                    <Divider dashed style={{ marginBlock: 4 }} />

                    <div>
                        <Switch
                            checked={config.useDateRange}
                            size="small"
                            onChange={(check) => setConfig((args) => ({
                                ...args,
                                useDateRange: check
                            }))}
                        /> Gunakan Range Tanggal
                    </div>
                    {config.useDateRange && <CheckOrderDateRange
                        value={[config.startDate, config.endDate]}
                        onChange={(range) => setConfig((args) => ({
                            ...args,
                            startDate: range[0],
                            endDate: range[1]
                        }))}
                    />}
                    <Divider dashed style={{ marginBlock: 4 }} />

                    <div>
                        <Switch
                            checked={config.useStatus}
                            size="small"
                            onChange={(check) => setConfig((args) => ({
                                ...args,
                                useStatus: check
                            }))}
                        /> Gunakan Status Custom
                    </div>
                    {config.useStatus && <CheckOrderSelectStatus
                        value={config.statusKeys}
                        onChange={(ids) => setConfig((args) => ({
                            ...args,
                            statusKeys: ids,
                        }))}
                    />}
                    <Divider dashed style={{ marginBlock: 4 }} />

                    <Button
                        onClick={() => {
                            let rename = name
                            if (rename) {
                                rename = rename + ".csv"
                            } else {
                                rename = "cekorder_tokopedia.csv"
                            }

                            props.onFinish(rename, config)
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
