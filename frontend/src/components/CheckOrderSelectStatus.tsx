import { Select, SelectProps } from "antd"
import React from "react"

import { DefaultOptionType } from "antd/es/select"
import orderStatus from "../assets/order_status.json"

type Props = SelectProps<string[], DefaultOptionType>

const CheckOrderStatus: React.FC<Props> = (props: Props) => {
    const options = orderStatus.map((opt) => ({
        label: opt.text,
        value: opt.key,
    }))

    return <Select
        mode="multiple"
        allowClear={false}
        options={options}
        placeholder="Pilih Status"
        {...props}
    />
}

export default CheckOrderStatus
