import { Alert, Button, Card, Input, message, Modal } from "antd"
import { TextAreaRef } from "antd/es/input/TextArea"
import React, { useState } from "react"
import { useRequest } from "../client"
import { BulkItem } from "../client/sdk_types"
import { FlexColumn } from "../styled_components"

export default function AddAccount(): React.ReactElement {
    const [modal, modalCtx] = Modal.useModal()
    const { sender } = useRequest("PostTokopediaAkunBulkAdd", {
        onError: console.log,
        onSuccess: () => {
            modal.confirm({
                title: "Input new account success",
                centered: true,
                content: "Clear textarea input?",
                icon: <i></i>,
                onOk: () => setAccountString(""),
            })
        },
    })

    const [accountString, setAccountString] = useState("")
    const textarea = React.createRef<TextAreaRef>()

    function bulkAddAction() {
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

            const errorLineContent =
                accountsList[(invalidFormat[0] as number) - 1]
            const [start, end] = [
                accountString.lastIndexOf(errorLineContent),
                accountString.lastIndexOf(errorLineContent) +
                    errorLineContent.length,
            ]

            textarea.current?.focus()
            if (textarea.current?.resizableTextArea?.textArea.selectionStart)
                textarea.current.resizableTextArea.textArea.selectionStart =
                    start
            if (textarea.current?.resizableTextArea?.textArea.selectionEnd)
                textarea.current.resizableTextArea.textArea.selectionEnd = end

            return
        }

        const payload = {
            data: data.filter(Boolean) as BulkItem[],
        }

        sender({
            method: "post",
            path: "tokopedia/akun/bulk_add",
            payload: payload,
        })
    }

    return (
        <FlexColumn>
            {modalCtx}
            <Card size="small"  title="Bulk Add Tokopedia Account">
                <FlexColumn>
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
                    <Button
                        style={{ boxShadow: "none" }}
                        type="primary"
                        onClick={bulkAddAction}
                    >
                        Add Account
                    </Button>
                </FlexColumn>
            </Card>
        </FlexColumn>
    )
}
