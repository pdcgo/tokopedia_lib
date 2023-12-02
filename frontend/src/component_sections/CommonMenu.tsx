import {
    ContainerOutlined,
    IdcardOutlined,
    RobotOutlined,
    UsergroupAddOutlined,
} from "@ant-design/icons"
import { Alert, Button, Card, Input, message } from "antd"
import { TextAreaRef } from "antd/es/input/TextArea"
import React, { Suspense, useState } from "react"
import { useRequest } from "../client"
import { useMutation } from "../client/sdk_mutation"
import {
    BulkItem,
    DriverAccount,
    VerifDriverAccount,
} from "../client/sdk_types"
import { Flex, FlexColumn } from "../styled_components"
import { accountPayloadChecker } from "../utils/accountPayloadChecker"
import { CheckOrderAkunItem } from "../client/newapisdk"

const CheckBotAsk = React.lazy(() => import("../components/CheckBotAsk"))
const CheckSubmitAsk = React.lazy(() => import("../components/CheckSubmitAsk"))
const CheckOrderAsk = React.lazy(() => import("../components/CheckOrderAsk"))

export default function CommonMenu() {
    const [showAsk, setShowAsk] = useState(false)
    const [showAskOrder, setShowAskOrder] = useState(false)
    const [showAskKtp, setShowAskKtp] = useState(false)
    const { sender } = useRequest("PostTokopediaAkunBulkAdd", {
        onError: (e) => message.error(`Error: ${e.msg}`),
        onSuccess: () => {
            message.success("Success!")
        },
    })
    const { sender: checkbot } = useRequest("PutTokopediaCekbotRun")
    const { sender: verifKtp } = useRequest("PutTokopediaCheckVerifRun")
    const { mutate: checkorder } = useMutation("PutTokopediaCekorderRun")

    const [accountString, setAccountString] = useState("")
    const textarea = React.createRef<TextAreaRef>()

    function bulkAddAction() {
        accountPayloadChecker(
            accountString,
            textarea,
            (warn) => {
                message.warning({ content: warn, key: "ghigggj" })
            },
            (data) => {
                const payload = {
                    data: data.filter(Boolean) as BulkItem[],
                }

                sender({
                    method: "post",
                    path: "tokopedia/akun/bulk_add",
                    payload: payload,
                })
            }
        )
    }

    function checkBotAction(filename: string) {
        accountPayloadChecker(
            accountString,
            textarea,
            (warn) => {
                message.warning({ content: warn, key: "ghigggj" })
            },
            (data) => {
                const payload = data.filter(Boolean) as DriverAccount[]

                checkbot({
                    method: "put",
                    path: "tokopedia/cekbot/run",
                    payload: {
                        fname: filename,
                        Akuns: payload,
                    },
                })
            }
        )
    }

    function checkSubmitAction(filename: string) {
        accountPayloadChecker(
            accountString,
            textarea,
            (warn) => {
                message.warning({ content: warn, key: "vfdvf" })
            },
            (data) => {
                const payload = data.filter(Boolean) as VerifDriverAccount[]

                verifKtp({
                    method: "put",
                    path: "tokopedia/check_verif/run",
                    payload: {
                        fname: filename,
                        Akuns: payload,
                    },
                })
            }
        )
    }

    function checkOrderAction(filename: string) {
        accountPayloadChecker(
            accountString,
            textarea,
            (warn) => {
                message.warning({ content: warn, key: "corder" })
            },
            (data) => {
                const payload = data.filter(Boolean) as CheckOrderAkunItem[]

                checkorder({
                    query: {
                        base: "",
                        fname: "",
                        output: filename,
                    },
                }, payload)
            }
        )
    }

    return (
        <Card size="small" title="Bulk Add Tokopedia Account">
            <FlexColumn>
                <Suspense fallback={<></>}>
                    <CheckBotAsk
                        onFinish={(name) => {
                            setShowAsk(false)
                            checkBotAction(name)
                        }}
                        open={showAsk}
                        onCancel={() => setShowAsk(false)}
                    />
                </Suspense>
                <Suspense fallback={<></>}>
                    <CheckSubmitAsk
                        onFinish={(name) => {
                            setShowAskKtp(false)
                            checkSubmitAction(name)
                        }}
                        open={showAskKtp}
                        onCancel={() => setShowAskKtp(false)}
                    />
                </Suspense>
                <Suspense fallback={<></>}>
                    <CheckOrderAsk
                        onFinish={(name) => {
                            setShowAskOrder(false)
                            checkOrderAction(name)
                        }}
                        open={showAskOrder}
                        onCancel={() => setShowAskOrder(false)}
                    />
                </Suspense>
                <Alert
                    type="info"
                    message="Format: username|password|otp_secret"
                    showIcon
                />
                <Input.TextArea
                    ref={textarea}
                    size="large"
                    autoSize={{ minRows: 24, maxRows: 24 }}
                    value={accountString}
                    onChange={(e) => setAccountString(e.target.value)}
                />
                <Flex>
                    <Button
                        style={{ boxShadow: "none" }}
                        type="primary"
                        onClick={bulkAddAction}
                        icon={<UsergroupAddOutlined rev="add-account" />}
                    >
                        Add Account
                    </Button>
                    <div style={{ flex: 1 }}></div>
                    <Button
                        style={{
                            backgroundColor: "#1677ff",
                            boxShadow: "none",
                            color: "#fff",
                        }}
                        type="primary"
                        icon={<ContainerOutlined rev="check-order" />}
                        onClick={() => setShowAskOrder(true)}
                    >
                        Check Order
                    </Button>
                    <Button
                        style={{
                            backgroundColor: "#005246",
                            boxShadow: "none",
                            color: "#fff",
                        }}
                        type="primary"
                        icon={<RobotOutlined rev="check-bot" />}
                        onClick={() => setShowAsk(true)}
                    >
                        Check Bot
                    </Button>
                    <Button
                        type="primary"
                        style={{
                            backgroundColor: "#C2418D",
                            boxShadow: "none",
                            color: "#fff",
                        }}
                        icon={<IdcardOutlined rev="id-card" />}
                        onClick={() => setShowAskKtp((f) => !f)}
                    >
                        Check Submit KTP
                    </Button>
                </Flex>
            </FlexColumn>
        </Card>
    )
}
