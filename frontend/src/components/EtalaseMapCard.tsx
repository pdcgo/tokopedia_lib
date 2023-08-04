/* eslint-disable react-hooks/exhaustive-deps */
import { Breadcrumb, Button, Card, Divider, Input, Select, message } from "antd"
import { useState } from "react"
import { SenderConfigs, UseQueryOptions, useRequest } from "../client"
import {
    EtalasePayload,
    Response,
    ShopeeEtalaseMapItem,
} from "../client/sdk_types"
import { Flex, FlexColumn } from "../styled_components"

type RFn = (
    config: SenderConfigs<
        "get",
        "tokopedia/etalase_map/list_etalase",
        undefined,
        undefined
    >,
    senderOptions?: UseQueryOptions<EtalasePayload[], Response> | undefined
) => Promise<unknown>

export type Props = {
    item: ShopeeEtalaseMapItem
    style?: React.CSSProperties
    etalases?: { value: string; label: string }[]
    refetchFn?: RFn
}

export default function EtalaseMapCard(props: Props) {
    const [newName, setNewName] = useState<string>()
    const [selected, setSelected] = useState<string | undefined>(
        props.item.EtalaseName || undefined
    )

    const { sender } = useRequest("PutTokopediaEtalaseMapUpdate")

    function onSelect(key: string) {
        if (!newName)
            sender(
                {
                    method: "put",
                    path: "tokopedia/etalase_map/update",
                    payload: [
                        {
                            category_id: props.item.tokpedia_id,
                            etalase_name: key,
                            ID: 0,
                        },
                    ],
                },
                {
                    onSuccess() {
                        setSelected(key)
                        message.success(`Updated as ${key}`)
                    },
                }
            )
    }

    function onAdd() {
        if (newName) {
            if (props.item.tokpedia_id) {
                sender(
                    {
                        method: "put",
                        path: "tokopedia/etalase_map/update",
                        payload: [
                            {
                                category_id: props.item.tokpedia_id,
                                etalase_name: newName,
                                ID: 0,
                            },
                        ],
                    },
                    {
                        onError(err) {
                            message.error(JSON.stringify(err))
                        },
                        onSuccess() {
                            message.success(`Updated as ${newName}`)
                            props.refetchFn?.(
                                {
                                    method: "get",
                                    path: "tokopedia/etalase_map/list_etalase",
                                },
                                {
                                    onSuccess() {
                                        setSelected(newName)
                                    },
                                }
                            )
                        },
                    }
                )
            }
        }
    }

    return (
        <Card size="small" type="inner" style={props.style}>
            <FlexColumn
                style={{ justifyContent: "space-between", height: "100%" }}
            >
                <Flex style={{ justifyContent: "space-between", flex: 1 }}>
                    <Breadcrumb
                        items={props.item.ShopeeCategoryName.map((f) => ({
                            title: f,
                        }))}
                        separator=">"
                    />
                    <div style={{ width: "20px" }}></div>
                    <span style={{ flexShrink: 0 }}>
                        <strong>{props.item.product_count || 0}</strong> Product
                    </span>
                </Flex>
                <Select
                    placeholder={
                        !props.item.tokpedia_id
                            ? "Check category mapping"
                            : "Choose etalase"
                    }
                    disabled={!props.item.tokpedia_id}
                    dropdownRender={(menu) => {
                        return (
                            <>
                                {menu}
                                <Divider style={{ marginBlock: 4 }}></Divider>
                                <Flex style={{ columnGap: 5 }}>
                                    <Input
                                        value={newName}
                                        onKeyDown={(e) => {
                                            if (e.code == "Enter") {
                                                onAdd()
                                            }
                                        }}
                                        onChange={(e) =>
                                            setNewName(e.target.value)
                                        }
                                        placeholder="Create new etalase..."
                                    />
                                    <Button
                                        disabled={!newName}
                                        onClick={onAdd}
                                        type="primary"
                                    >
                                        Add
                                    </Button>
                                </Flex>
                            </>
                        )
                    }}
                    value={selected}
                    onChange={(key) => onSelect(key)}
                    onClick={() => setNewName(undefined)}
                >
                    {props.etalases
                        ?.sort((a, v) =>
                            a.label.toLowerCase() < v.label.toLowerCase()
                                ? -1
                                : 0
                        )
                        .map((e) => (
                            <Select.Option key={e.label}>
                                {e.label}
                            </Select.Option>
                        ))}
                </Select>
            </FlexColumn>
        </Card>
    )
}
