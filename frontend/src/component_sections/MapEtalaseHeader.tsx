/* eslint-disable react-hooks/exhaustive-deps */
import React from "react"
import LabelInput from "../components/LabelInput"
import { Button, Card, Select, Typography } from "antd"
import { Flex } from "../styled_components"
import { useRequest } from "../client"

export type Props = {
    collection?: string
    onChangeCollection?: (name: string) => void
}

export default function MapEtalaseHeader(props: Props) {
    const { sender, response, pending } = useRequest("GetV1ProductNamespaceAll")

    const namespaces = React.useMemo(() => {
        if (!response) {
            return []
        }

        return response.map((r) => ({ value: r.name, label: r.name }))
    }, [response])

    React.useEffect(() => {
        const c = new AbortController()
        const signal = c.signal

        sender(
            { method: "get", path: "v1/product/namespace_all" },
            {
                signal,
                onSuccess(data) {
                    const f = data[0]
                    if (f) {
                        props.onChangeCollection?.(f.name)
                    }
                },
            }
        )

        return () => {
            c.abort()
            props.onChangeCollection?.("")
        }
    }, [])

    return (
        <Card
            size="small"
            title={
                <Typography.Text>Map Etalase From Collections</Typography.Text>
            }
        >
            <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "flex-end",
                }}
            >
                <LabelInput label="Collections :">
                    <Select
                        style={{ width: "300px" }}
                        placeholder="Choose collection"
                        options={namespaces}
                        loading={pending}
                        value={props.collection}
                        onChange={props.onChangeCollection}
                    ></Select>
                </LabelInput>
                <Button>Reset</Button>
            </Flex>
        </Card>
    )
}
