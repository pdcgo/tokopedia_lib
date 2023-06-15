/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Card, Select, Typography, message } from "antd"
import { useEffect, useState } from "react"
import { useRequest } from "../client"
import { CategoryAllListLiteRes } from "../client/sdk_types"
import { ListMapper, ListMapperActions } from "../store/listMapper"
import { Flex, FlexColumn } from "../styled_components"

type SelectOption = { label: string; value: string }

export type MapperHeaderProps = {
    namespace: string | null
    namespaces: SelectOption[]
    onChangeNamespace: React.Dispatch<React.SetStateAction<string | null>>

    list: ListMapper[]
    initEffect: ListMapperActions["initEffect"]
    listCategoryTokopedia: CategoryAllListLiteRes | null
}

export default function MapperHeader(props: MapperHeaderProps) {
    const [gettingSuggest, setGettingSuggest] = useState(false)
    const { sender: runAutoSuggets } = useRequest(
        "PutTokopediaMapperAutosuggest"
    )
    const { sender: autoSuggestChecker } = useRequest(
        "GetTokopediaMapperAutosuggest"
    )

    const { sender: saveMapSender, pending: saveMapPending } = useRequest(
        "PutTokopediaMapperMap"
    )

    const mapSaver = () => {
        saveMapSender({
            method: "put",
            path: "tokopedia/mapper/map",
            payload: props.list.map((payload) => ({
                shopee_id: payload.shopeeCategoryId,
                tokopedia_id:
                    payload.tokopediaCategoryIds.map(Number)[
                        payload.tokopediaCategoryIds.length - 1
                    ],
            })),
        })
    }

    const reset = () => {
        saveMapSender(
            {
                method: "put",
                path: "tokopedia/mapper/map",
                payload: props.list.map((payload) => ({
                    shopee_id: payload.shopeeCategoryId,
                    tokopedia_id: 0,
                })),
            },
            {
                onSuccess() {
                    message.success("Update data success!")
                },
                onError() {
                    message.error("Update data error!")
                },
            }
        )
        if (props.namespace) {
            props.initEffect(props.namespace, props.listCategoryTokopedia)
        }
    }

    const getSuggest = () => {
        if (!gettingSuggest && props.namespace) {
            setGettingSuggest(true)
            runAutoSuggets({
                method: "put",
                path: "tokopedia/mapper/autosuggest",
                params: { collection: props.namespace },
            })
        }
    }

    useEffect(() => {
        const int: any = setInterval(() => {
            if (!gettingSuggest) {
                if (props.namespace && !gettingSuggest) {
                    props.initEffect(
                        props.namespace,
                        props.listCategoryTokopedia
                    )
                }
                return clearInterval(int)
            }
            autoSuggestChecker(
                {
                    method: "get",
                    path: "tokopedia/mapper/autosuggest",
                },
                {
                    onSuccess: (data) => {
                        if (data?.status === "STOPPED") setGettingSuggest(false)
                    },
                    onError: (e) => {
                        setGettingSuggest(false)
                        console.log(e)
                    },
                }
            )
        }, 1000)

        return () => clearInterval(int)
    }, [gettingSuggest])

    return (
        <Card
            size="small"
            title={
                <Typography.Text>
                    Map Category From Shopee to Tokopedia
                </Typography.Text>
            }
        >
            <Flex
                style={{
                    justifyContent: "space-between",
                    alignItems: "end",
                }}
            >
                <FlexColumn style={{ rowGap: "5px" }}>
                    <Typography.Text>Collections :</Typography.Text>
                    <Select
                        style={{ width: "300px" }}
                        placeholder="Choose Collection"
                        value={props.namespace}
                        onChange={props.onChangeNamespace}
                        options={props.namespaces}
                    />
                </FlexColumn>
                <Flex style={{ rowGap: "5px", justifyContent: "flex-end" }}>
                    <Button
                        loading={gettingSuggest}
                        disabled={gettingSuggest}
                        onClick={getSuggest}
                    >
                        Use Suggest
                    </Button>
                    <Button
                        style={{ backgroundColor: "#005246" }}
                        type="primary"
                        onClick={reset}
                    >
                        Reset All
                    </Button>
                    <Button
                        loading={saveMapPending}
                        onClick={mapSaver}
                        type="primary"
                    >
                        Save Mapping
                    </Button>
                </Flex>
            </Flex>
        </Card>
    )
}
