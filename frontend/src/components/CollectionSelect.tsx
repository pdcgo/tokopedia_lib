import { Select, SelectProps } from "antd"
import React from "react"

import { Selection } from "../store/listProfile"
import { Mode } from "../component_sections/UploadHeader"

export interface CollectionSelectProps extends Omit<SelectProps<string>, "options" | "mode"> {
    mode: Mode
    collections: Selection[]
    manualCollections: Selection[]
}

const CollectionSelect: React.FC<CollectionSelectProps> = (props: CollectionSelectProps) => {
    const { mode, collections, manualCollections, value, ...selecProps } = props

    const options = mode === "tokopedia_manual"
        ? manualCollections
        : collections
    
    console.log(mode)

    const findvalue = options.find((opt) => opt.label === value)

    return <Select
        value={findvalue?.value}
        options={[
            {
                value: "",
                label: "Choose Collection",
                disabled: true,
            },
            ...options,
        ]}
        placeholder="Choose Collection Data"
        {...selecProps}
    />
}

export default CollectionSelect