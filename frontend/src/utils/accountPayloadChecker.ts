import { TextAreaRef } from "antd/es/input/TextArea"

export
    const accountPayloadChecker = (
        accountString: string,
        textarea: React.RefObject<TextAreaRef>,
        onWarning?: (warn: string) => void,
        onPass?: (data: ({
            password: string;
            secret: string;
            username: string;
        } | null)[]) => void
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
            onWarning?.(`Invalid format on line: ${invalidFormat[0]}`)

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

            return
        }

        onPass?.(data)
    }