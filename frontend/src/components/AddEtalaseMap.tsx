/* eslint-disable react-hooks/exhaustive-deps */
import {
    Button,
    Card,
    Divider,
    Input,
    Modal,
    ModalProps,
    Select,
    message,
} from "antd"
import { Flex, FlexColumn } from "../styled_components"
import { useRequest } from "../client"
import LabelInput from "./LabelInput"
import React from "react"
import { unwrapDeepCategory } from "../utils/unwrapDeepCategory"

type ExtendedProps = {
    onFinish?: () => void
}

export default function AddEtalaseMap(props: ModalProps & ExtendedProps) {
    const {
        sender,
        response: catListTokopedia,
        pending,
    } = useRequest("GetTokopediaCategoryList")
    const { sender: poster, pending: pendingPoster } = useRequest(
        "PostTokopediaEtalaseMapAdd"
    )
    const [selected, setSelected] = React.useState<string[]>([])
    const [name, setName] = React.useState<string | undefined>()

    const options = React.useMemo(() => {
        if (catListTokopedia?.data.categoryAllListLite) {
            return unwrapDeepCategory(catListTokopedia)
        }

        return []
    }, [catListTokopedia])

    const onChange = (value: string[]) => {
        setSelected(value)
    }

    const onSave = () => {
        if (name && selected.length) {
            const cat_ids = selected
                .map((sel) => {
                    for (const opt of options) {
                        if (opt.label == sel) return opt.value
                    }

                    return
                })
                .filter((v) => typeof v == "number") as number[]

            poster(
                {
                    method: "post",
                    path: "tokopedia/etalase_map/add",
                    payload: {
                        etalase: name,
                        cat_ids: cat_ids,
                    },
                },
                {
                    onSuccess: () => {
                        props.onFinish?.()
                    },
                }
            )
        } else {
            message.warning({
                content: "Payload invalid!",
                duration: 3,
            })
        }
    }

    React.useEffect(() => {
        if (props.open) {
            sender({ method: "get", path: "tokopedia/category/list" })
        } else {
            setSelected([])
            setName(undefined)
        }
    }, [props.open])

    return (
        <Modal
            maskClosable={false}
            width={530}
            footer={false}
            closable={false}
            centered
            {...props}
        >
            <Card title="Etalase Map Fields" size="small" type="inner">
                <FlexColumn>
                    <Flex
                        style={{
                            maxHeight: 220,
                            alignItems: "flex-start",
                            overflowY: "auto",
                            paddingRight: 3,
                        }}
                    >
                        <LabelInput
                            style={{ width: "100%" }}
                            label="Categories :"
                        >
                            <Select
                                loading={pending}
                                mode="multiple"
                                options={options.map((o) => ({
                                    value: o.label,
                                    label: o.label,
                                }))}
                                placeholder="Select category for etalase"
                                onChange={onChange}
                                value={selected}
                                allowClear
                            />
                        </LabelInput>
                    </Flex>
                    <LabelInput
                        style={{ width: "100%" }}
                        label="Etalase Name :"
                    >
                        <Input
                            onChange={(e) => setName(e.target.value)}
                            value={name}
                            placeholder="Pakaian Wanita"
                        />
                    </LabelInput>
                    <Divider style={{ marginBlock: 3 }} dashed />
                    <Button
                        loading={pendingPoster}
                        onClick={onSave}
                        block={false}
                        type="primary"
                    >
                        Save
                    </Button>
                    <Button onClick={props.onFinish} block={false}>
                        Cancel
                    </Button>
                </FlexColumn>
            </Card>
        </Modal>
    )
}
