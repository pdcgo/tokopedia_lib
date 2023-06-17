import { Card, Divider, Radio } from "antd"
import React, { Suspense, useState } from "react"
import { FlexColumn } from "../styled_components"

const CommonMenu = React.lazy(() => import("../component_sections/CommonMenu"))
const Deleter = React.lazy(() => import("../component_sections/AccountDeleter"))

export default function AddAccount(): React.ReactElement {
    const [mode, setMode] = useState("common_menu")

    return (
        <FlexColumn>
            <Radio.Group value={mode} onChange={(v) => setMode(v.target.value)}>
                <Radio.Button value="common_menu">Common Menu</Radio.Button>
                <Radio.Button value="deleter">
                    Tokopedia Product Deleter
                </Radio.Button>
            </Radio.Group>
            <Divider dashed style={{ marginBlock: 2 }} />
            {mode == "common_menu" && (
                <Suspense fallback={<Card loading />}>
                    <CommonMenu />
                </Suspense>
            )}
            {mode == "deleter" && (
                <Suspense fallback={<Card loading />}>
                    <Deleter />
                </Suspense>
            )}
        </FlexColumn>
    )
}
