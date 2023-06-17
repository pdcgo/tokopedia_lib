import { Typography } from "antd"
import { FlexColumn } from "../styled_components"

export type LabelInput = {
    label?: string
    children?: React.ReactElement
    style?: React.CSSProperties
}

export default function LabelInput(props: LabelInput) {
    return (
        <FlexColumn style={{ rowGap: "5px", ...props.style }}>
            <Typography.Text>{props.label}</Typography.Text>
            {props.children}
        </FlexColumn>
    )
}
