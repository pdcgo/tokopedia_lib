import { RobotOutlined } from "@ant-design/icons"
import {
    Alert,
    Button,
    Card,
    DatePicker,
    Input,
    InputNumber,
    Select,
    message
} from "antd"
import { TextAreaRef } from "antd/es/input/TextArea"
import dayjs from "dayjs"
import React, { useState } from "react"
import { useRequest } from "../client"
import { AkunDeleteItem } from "../client/sdk_types"
import LabelInput from "../components/LabelInput"
import { Flex, FlexColumn } from "../styled_components"
import { accountPayloadChecker } from "../utils/accountPayloadChecker"

export default function AccountDeleter() {
    const { sender: filterPutter } = useRequest("PutTokopediaDeleterSetting")
    const { sender: runner } = useRequest("PutTokopediaDeleterRunDelete")
    const [accounts, setAccounts] = useState("")
    const textarea = React.createRef<TextAreaRef>()

    // bagian setting filter
    const [limitCon, setLimitCon] = useState(0)
    const [limitDel, setLimitDel] = useState(0)
    const [status, setStatus] = useState("")
    const [keywords, setKeywords] = useState("")
    const [dateRange, setDateRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([
        dayjs().subtract(30, "day"),
        dayjs(),
    ])
    const [view, setView] = useState<[number | null, number | null]>([
        null,
        null,
    ])
    const [sold, setSold] = useState<[number | null, number | null]>([
        null,
        null,
    ])

    const deleteAction = () => {
        accountPayloadChecker(
            accounts,
            textarea,
            (warn) => {
                message.warning({
                    content: warn,
                    key: "warningvalidation",
                })
            },
            (data) => {
                const payload = data.filter(Boolean) as AkunDeleteItem[]

                filterPutter(
                    {
                        method: "put",
                        path: "tokopedia/deleter/setting",
                        payload: {
                            category_id: "",
                            akuns: payload,
                            start_time: dateRange[0].unix(),
                            end_time: dateRange[1].unix(),
                            limit_concurent: limitCon,
                            limit_product: limitDel,
                            product_status: status,
                            title: keywords
                                .split("\n")
                                .map((line) => line.trim())
                                .filter(Boolean),
                            ...(sold[0] != null && sold[1] != null
                                ? {
                                      sold_filter: {
                                          max: sold[1],
                                          min: sold[0],
                                      },
                                  }
                                : {}),
                            ...(view[0] != null && view[1] != null
                                ? {
                                      view_filter: {
                                          max: view[1],
                                          min: view[0],
                                      },
                                  }
                                : {}),
                        },
                    },
                    {
                        onSuccess() {
                            message.info("Running deleter...")
                            runner(
                                {
                                    method: "put",
                                    path: "tokopedia/deleter/run_delete",
                                },
                                {
                                    onError(err) {
                                        message.error(err.msg)
                                    },
                                }
                            )
                        },
                        onError(e) {
                            message.error(e.msg)
                        },
                    }
                )
            }
        )
    }

    return (
        <Card size="small" title="Tokopedia Product Deleter">
            <FlexColumn>
                <Alert
                    type="info"
                    message="Format: username|password|otp_secret"
                    showIcon
                />
                <FlexColumn>
                    <Input.TextArea
                        size="large"
                        autoSize={{ minRows: 24, maxRows: 24 }}
                        value={accounts}
                        onChange={(e) => setAccounts(e.target.value)}
                        ref={textarea}
                    />
                    <Card size="small" type="inner" title="Filter">
                        <Flex style={{ alignItems: "start" }}>
                            <LabelInput
                                style={{ flex: 1 }}
                                label="Filter Title (Regex Allowed) :"
                            >
                                <Input.TextArea
                                    size="large"
                                    autoSize={{ minRows: 13, maxRows: 13 }}
                                    placeholder={`rokok herbal
obat pelangsing
regex-->obat|jamu|ramuan
                                    `}
                                    value={keywords}
                                    onChange={(e) =>
                                        setKeywords(e.target.value)
                                    }
                                />
                            </LabelInput>
                            <FlexColumn style={{ width: "50%" }}>
                                <Flex style={{ width: "100%" }}>
                                    <LabelInput
                                        style={{ flex: 1 }}
                                        label="Limit Concurrent :"
                                    >
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="10"
                                            value={limitCon}
                                            onChange={(e) =>
                                                e && setLimitCon(e)
                                            }
                                        />
                                    </LabelInput>
                                    <LabelInput
                                        style={{ flex: 1 }}
                                        label="Limit Delete :"
                                    >
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="10"
                                            value={limitDel}
                                            onChange={(e) =>
                                                e && setLimitDel(e)
                                            }
                                        />
                                    </LabelInput>
                                </Flex>
                                <Flex style={{ width: "100%" }}>
                                    <LabelInput
                                        style={{ flex: 1 }}
                                        label="Sold :"
                                    >
                                        <Flex
                                            style={{
                                                alignItems: "center",
                                                columnGap: 7,
                                            }}
                                        >
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Min"
                                                value={sold[0] || null}
                                                onChange={(e) =>
                                                    setSold((s) => [e, s[1]])
                                                }
                                            />
                                            -
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Max"
                                                value={sold[1] || null}
                                                onChange={(e) =>
                                                    setSold((s) => [s[0], e])
                                                }
                                            />
                                        </Flex>
                                    </LabelInput>
                                    <LabelInput
                                        style={{ flex: 1 }}
                                        label="View :"
                                    >
                                        <Flex
                                            style={{
                                                alignItems: "center",
                                                columnGap: 7,
                                            }}
                                        >
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Min"
                                                value={view[0] || null}
                                                onChange={(e) =>
                                                    setView((v) => [e, v[1]])
                                                }
                                            />
                                            -
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Max"
                                                value={view[1] || null}
                                                onChange={(e) =>
                                                    setView((v) => [v[0], e])
                                                }
                                            />
                                        </Flex>
                                    </LabelInput>
                                </Flex>
                                <LabelInput label="Product Status :">
                                    <Select
                                        value={status}
                                        onChange={setStatus}
                                        defaultValue=""
                                    >
                                        <Select.Option value="">
                                            All (No Status)
                                        </Select.Option>
                                        <Select.Option value="ACTIVE">
                                            Active
                                        </Select.Option>
                                        <Select.Option value="VIOLATION">
                                            Pelanggaran
                                        </Select.Option>
                                        <Select.Option value="INACTIVE">
                                            InActive
                                        </Select.Option>
                                    </Select>
                                </LabelInput>
                                <LabelInput
                                    style={{ flex: 2 }}
                                    label="Date Range :"
                                >
                                    <DatePicker.RangePicker
                                        onChange={(d) => {
                                            if (d?.[0] && d?.[1])
                                                setDateRange([d?.[0], d?.[1]])
                                        }}
                                        format="DD MMMM YYYY"
                                        value={dateRange}
                                    />
                                </LabelInput>
                                <Button
                                    style={{
                                        backgroundColor: "#C2418D",
                                        boxShadow: "none",
                                        color: "#fff",
                                    }}
                                    type="primary"
                                    onClick={deleteAction}
                                    icon={<RobotOutlined rev="run-delete" />}
                                >
                                    Run Delete
                                </Button>
                            </FlexColumn>
                        </Flex>
                    </Card>
                </FlexColumn>
            </FlexColumn>
        </Card>
    )
}
