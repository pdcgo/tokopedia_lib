import React, { Suspense, useState } from "react"
import { RobotOutlined } from "@ant-design/icons"
import { Card, Alert, Input, Button, message } from "antd"
import { FlexColumn, Flex } from "../styled_components"
import { TextAreaRef } from "antd/es/input/TextArea"
import { useRequest } from "../client"
import { BulkItem, DriverAccount } from "../client/sdk_types"

const CheckBotAsk = React.lazy(() => import("../components/CheckBotAsk"))

const accountPayloadChecker = (
    accountString: string,
    textarea: React.RefObject<TextAreaRef>
) => {
    const accountsList = accountString.split("\n")
    const data = accountsList.map((account) => {
        const [username, password, secret] = account.split("|")

        if (!username || !password || !secret) return null

        return {
            password: password.trim(),
            secret: secret.trim(),
            username: username.trim().toLowerCase(),
        }
    })

    const invalidFormat = data
        .map((d, i) => (d == null ? i + 1 : null))
        .filter((c) => c !== null)

    if (invalidFormat.length) {
        message.error({
            content: (
                <span>
                    Invalid format on line:{" "}
                    <i>
                        <strong>{invalidFormat[0]}</strong>
                    </i>
                </span>
            ),
        })

        const errorLineContent = accountsList[(invalidFormat[0] as number) - 1]
        const [start, end] = [
            accountString.lastIndexOf(errorLineContent),
            accountString.lastIndexOf(errorLineContent) +
                errorLineContent.length,
        ]

        textarea.current?.focus()
        if (textarea.current?.resizableTextArea?.textArea.selectionStart)
            textarea.current.resizableTextArea.textArea.selectionStart = start
        if (textarea.current?.resizableTextArea?.textArea.selectionEnd)
            textarea.current.resizableTextArea.textArea.selectionEnd = end

        return null
    }

    return data
}

export default function CommonMenu() {
    const [showAsk, setShowAsk] = useState(false)
    const { sender } = useRequest("PostTokopediaAkunBulkAdd", {
        onError: (e) => message.error(`Error: ${e.msg}`),
        onSuccess: () => {
            message.success("Success!")
        },
    })
    const { sender: checkbot } = useRequest("PutTokopediaCekbotRun")

    const [accountString, setAccountString] = useState("")
    const textarea = React.createRef<TextAreaRef>()

    function bulkAddAction() {
        const data = accountPayloadChecker(accountString, textarea)

        if (data) {
            const payload = {
                data: data.filter(Boolean) as BulkItem[],
            }

            sender({
                method: "post",
                path: "tokopedia/akun/bulk_add",
                payload: payload,
            })
        }
    }

    function checkBotAction(filename: string) {
        const data = accountPayloadChecker(accountString, textarea)
        if (data) {
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
                    >
                        Add Account
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
                </Flex>
            </FlexColumn>
        </Card>
    )
}
