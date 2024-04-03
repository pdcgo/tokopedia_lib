import {
    Button,
    Card,
    DatePicker,
    Divider,
    Input,
    InputNumber,
    Modal,
    ModalProps,
    Select,
    TimeRangePickerProps,
    Typography
} from "antd"
import { useState } from "react"
import { FlexColumn } from "../styled_components"
// import CheckOrderDateRange from "./CheckOrderDateRange"
// import CheckOrderSelectStatus from "./CheckOrderSelectStatus"
import { TokopediaTokoLiburConfig } from "../client/newapisdk"
import dayjs from "dayjs"

export type TokoLiburProps = {
    onFinish: (config: TokopediaTokoLiburConfig) => void
    onCancel: () => void
}

const DateFormat = "DD/MM/YYYY"

const now = dayjs()
const rangePresets: TimeRangePickerProps["presets"] = [
    { label: "1 Minggu", value: [now, now.add(7, "day")] },
    { label: "2 Minggu", value: [now, now.add(14, "day")] },
    { label: "1 Bulan", value: [now, dayjs().add(1, "month")] },
    { label: "6 Bulan", value: [now, dayjs().add(6, "month")] },
    { label: "1 Tahun", value: [now, dayjs().add(1, "year")] },
]

export default function TokoLiburAsk(props: TokoLiburProps & ModalProps) {
    const [query, setQuery] = useState<TokopediaTokoLiburConfig>({
        libur: true,
        report: "tokopedia_toko_libur",
        limit: 3,
        closeStart: now.unix(),
        closeEnd: now.add(7, "day").unix()
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
                        value={query.report}
                        onChange={(e) => setQuery({ ...query, report: e.target.value })}
                        addonAfter=".csv"
                        placeholder="Boleh dikosongi"
                    />
                    <Typography.Text>
                        Default filename: tokopedia_toko_libur.csv
                    </Typography.Text>
                    <Divider dashed style={{ marginBlock: 4 }} />

                    <Select
                        value={query.libur}
                        options={[
                            { value: true, label: "Liburkan Toko" },
                            { value: false, label: "Buka Toko Libur" },
                        ]}
                        onChange={(libur) => setQuery({ ...query, libur })}
                    />

                    <InputNumber
                        value={query.limit}
                        addonBefore="Concurrent"
                        onChange={(v) => setQuery({ ...query, limit: v || 0 })}
                    />

                    <DatePicker.RangePicker
                        allowClear={false}
                        value={[
                            dayjs.unix(query.closeStart),
                            dayjs.unix(query.closeEnd),
                        ]}
                        disabled={!query.libur}
                        disabledDate={(current) => current && current > dayjs().endOf("day")}
                        presets={rangePresets}
                        format={DateFormat}
                        onChange={(ranges) => {
                            setQuery({
                                ...query,
                                closeStart: ranges?.[0]?.unix() || 0,
                                closeEnd: ranges?.[1]?.unix() || 0,
                            })
                        }}
                    />
                    <Divider dashed style={{ marginBlock: 4 }} />

                    <Button
                        onClick={() => {
                            const requery = { ...query }
                            if (requery.report) {
                                requery.report = requery.report + ".csv"
                            } else {
                                requery.report = "tokopedia_toko_libur.csv"
                            }
                            props.onFinish(requery)
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
