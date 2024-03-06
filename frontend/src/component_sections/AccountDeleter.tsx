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
import { AxiosError } from "axios"
import dayjs from "dayjs"
import React, { useState } from "react"

import { AkunDeleteItem, TokopediaDeleteConfig } from "../client/newapisdk"
import { useMutation } from "../client/sdk_mutation"
import { Response } from "../client/sdk_types"
import LabelInput from "../components/LabelInput"
import { Flex, FlexColumn } from "../styled_components"
import { accountPayloadChecker } from "../utils/accountPayloadChecker"

export default function AccountDeleter() {
    const { mutate: filterPutter } = useMutation("PutTokopediaDeleterSetting")
    const { mutate: runner } = useMutation("PutTokopediaDeleterRunDelete")
    const [accounts, setAccounts] = useState("")
    const textarea = React.createRef<TextAreaRef>()

    // bagian setting filter
    const [config, setConfig] = useState<TokopediaDeleteConfig>({
        limit_concurent: 0,
        limit_product: 0,
        title: [],
        product_status: "",
        category_id: "",
        start_time: 0,
        end_time: 0,
        akuns: [],
        sold_filter: undefined,
        view_filter: undefined,
        price_filter: undefined
    })
    const [dateRange, setDateRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([
        dayjs().subtract(30, "day"),
        dayjs(),
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
                config.akuns = data.filter(Boolean) as AkunDeleteItem[]
                config.start_time = dateRange[0].unix()
                config.end_time = dateRange[1].unix()

                if (!config.sold_filter?.max) {
                    config.sold_filter = undefined
                }
                if (!config.view_filter?.max) {
                    config.view_filter = undefined
                }
                if (!config.price_filter?.max) {
                    config.price_filter = undefined
                }

                filterPutter({
                    onSuccess() {
                        message.info("Running deleter...")
                        runner({
                            onError(err) {
                                message.error((err as AxiosError<Response>).response?.data.msg)
                            },
                        })
                    },
                    onError(err) {
                        message.error((err as AxiosError<Response>).response?.data.msg)
                    },
                }, config)
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
                            <LabelInput style={{ flex: 1 }} label="Filter Title (Regex Allowed) :">
                                <Input.TextArea
                                    size="large"
                                    autoSize={{ minRows: 16, maxRows: 16 }}
                                    placeholder={`rokok herbal
obat pelangsing
regex-->obat|jamu|ramuan
                                    `}
                                    value={config.title.join("\n")}
                                    onChange={(e) => setConfig({
                                        ...config,
                                        title: e.target.value.split("\n")
                                            .map((line) => line.trim())
                                            .filter(Boolean)
                                    })}
                                />
                            </LabelInput>
                            <FlexColumn style={{ width: "50%" }}>
                                <Flex style={{ width: "100%" }}>
                                    <LabelInput style={{ flex: 1 }} label="Limit Concurrent :">
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="10"
                                            value={config.limit_concurent}
                                            onChange={(limit_concurent) => limit_concurent &&
                                                setConfig({ ...config, limit_concurent })
                                            }
                                        />
                                    </LabelInput>
                                    <LabelInput style={{ flex: 1 }} label="Limit Delete :">
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="10"
                                            value={config.limit_product}
                                            onChange={(limit_product) => limit_product &&
                                                setConfig({ ...config, limit_product })
                                            }
                                        />
                                    </LabelInput>
                                </Flex>
                                <Flex style={{ width: "100%" }}>
                                    <LabelInput style={{ flex: 1 }} label="Sold :">
                                        <Flex style={{
                                            alignItems: "center",
                                            columnGap: 7,
                                        }}>
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Min"
                                                value={config.sold_filter?.min}
                                                onChange={(e) => setConfig({
                                                    ...config,
                                                    sold_filter: {
                                                        min: e || 0,
                                                        max: config.sold_filter?.max || 0,
                                                    },
                                                })}
                                            />
                                            -
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Max"
                                                value={config.sold_filter?.max}
                                                onChange={(e) => setConfig({
                                                    ...config,
                                                    sold_filter: {
                                                        min: config.sold_filter?.min || 0,
                                                        max: e || 0,
                                                    },
                                                })}
                                            />
                                        </Flex>
                                    </LabelInput>
                                    <LabelInput style={{ flex: 1 }} label="View :">
                                        <Flex style={{
                                            alignItems: "center",
                                            columnGap: 7,
                                        }}>
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Min"
                                                value={config.view_filter?.min}
                                                onChange={(e) => setConfig({
                                                    ...config,
                                                    view_filter: {
                                                        min: e || 0,
                                                        max: config.view_filter?.max || 0,
                                                    },
                                                })}
                                            />
                                            -
                                            <InputNumber
                                                style={{ width: "100%" }}
                                                placeholder="Max"
                                                value={config.view_filter?.max}
                                                onChange={(e) => setConfig({
                                                    ...config,
                                                    view_filter: {
                                                        min: config.view_filter?.min || 0,
                                                        max: e || 0,
                                                    },
                                                })}
                                            />
                                        </Flex>
                                    </LabelInput>
                                </Flex>
                                <LabelInput style={{ flex: 1 }} label="Price :">
                                    <Flex style={{
                                        alignItems: "center",
                                        columnGap: 7,
                                    }}>
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="Min"
                                            prefix="Rp."
                                            value={config.price_filter?.min}
                                            onChange={(e) => setConfig({
                                                ...config,
                                                price_filter: {
                                                    min: e || 0,
                                                    max: config.price_filter?.max || 0,
                                                },
                                            })}
                                        />
                                        -
                                        <InputNumber
                                            style={{ width: "100%" }}
                                            placeholder="Max"
                                            prefix="Rp."
                                            value={config.price_filter?.max}
                                            onChange={(e) => setConfig({
                                                ...config,
                                                price_filter: {
                                                    min: config.price_filter?.min || 0,
                                                    max: e || 0,
                                                },
                                            })}
                                        />
                                    </Flex>
                                </LabelInput>
                                <LabelInput label="Product Status :">
                                    <Select
                                        value={config.product_status}
                                        onChange={(product_status) => setConfig({
                                            ...config,
                                            product_status,
                                        })}
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
                                <LabelInput style={{ flex: 2 }} label="Date Range :">
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
