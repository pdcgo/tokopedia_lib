import { DatePicker, TimeRangePickerProps } from "antd"
import { RangePickerDateProps } from "antd/es/date-picker/generatePicker"
import type { Dayjs } from "dayjs"
import dayjs from "dayjs"
import React from "react"

type RPProps = RangePickerDateProps<dayjs.Dayjs>

interface Props extends Omit<RPProps, "value" | "onChange"> {
    value: string[]
    onChange: (value: string[]) => void
}

const { RangePicker } = DatePicker
const DateFormat = "DD/MM/YYYY"

const rangePresets: TimeRangePickerProps["presets"] = [
    { label: "1 Minggu", value: [dayjs().add(-7, "day"), dayjs()] },
    { label: "2 Minggu", value: [dayjs().add(-14, "day"), dayjs()] },
    { label: "1 Bulan", value: [dayjs().add(-1, "month"), dayjs()] },
    { label: "6 Bulan", value: [dayjs().add(-6, "month"), dayjs()] },
    { label: "1 Tahun", value: [dayjs().add(-1, "year"), dayjs()] },
]

const end = dayjs()
const start = dayjs().add(-1, "month")

const CheckOrderDateRange: React.FC<Props> = (props: Props) => {
    const { value, onChange, ...rangeProps } = props


    const startDate = value[0] || start
    const endDate = value[1] || end

    function onRangeChange(_: null | (Dayjs | null)[], dateStrings: string[]) {
        onChange(dateStrings)
    }

    return <RangePicker
        allowClear={false}
        value={[
            dayjs(startDate, DateFormat),
            dayjs(endDate, DateFormat),
        ]}
        {...rangeProps}
        disabledDate={(current) => current && current > dayjs().endOf("day")}
        presets={rangePresets}
        format={DateFormat}
        onChange={onRangeChange}
    />
}

export default CheckOrderDateRange
