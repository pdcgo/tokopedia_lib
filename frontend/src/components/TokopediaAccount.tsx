import {
    Button,
    Card,
    Divider,
    Input,
    Modal,
    ModalProps,
    Typography,
} from "antd"
import { FlexColumn } from "../styled_components"
import { useState } from "react"
import { useRequest } from "../client"

export default function TokopediaAccount(
    props: ModalProps & { onFinish: () => void }
) {
    const [accountTarget, setAccountTarget] = useState({
        email: "",
        password: "",
        otp: "",
    })

    const { sender } = useRequest("PutTokopediaCategoryUpdateCategory", {
        onSuccess() {
            props.onFinish()
        },
    })

    function onOk() {
        sender({
            method: "put",
            path: "tokopedia/category/update_category",
            payload: {
                password: accountTarget.password,
                secret: accountTarget.otp,
                username: accountTarget.email,
            },
        })
    }

    return (
        <Modal
            {...props}
            onOk={onOk}
            centered
            closeIcon
            width={400}
            footer={null}
        >
            <Card
                size="small"
                type="inner"
                title="Tokopedia Account for Getting Cat List"
                bordered={false}
            >
                <FlexColumn>
                    <FlexColumn style={{ rowGap: 5 }}>
                        <Typography.Text>Email:</Typography.Text>
                        <Input
                            value={accountTarget.email}
                            onChange={(e) =>
                                setAccountTarget((a) => ({
                                    ...a,
                                    email: e.target.value,
                                }))
                            }
                            placeholder="silverrayleigh@yahoo.com"
                        />
                    </FlexColumn>
                    <FlexColumn style={{ rowGap: 5 }}>
                        <Typography.Text>Password:</Typography.Text>
                        <Input.Password
                            value={accountTarget.password}
                            onChange={(e) =>
                                setAccountTarget((a) => ({
                                    ...a,
                                    password: e.target.value,
                                }))
                            }
                            placeholder="⁎⁎⁎⁎⁎⁎⁎⁎"
                        />
                    </FlexColumn>
                    <FlexColumn style={{ rowGap: 5 }}>
                        <Typography.Text>OTP Secret:</Typography.Text>
                        <Input
                            value={accountTarget.otp}
                            onChange={(e) =>
                                setAccountTarget((a) => ({
                                    ...a,
                                    otp: e.target.value,
                                }))
                            }
                            placeholder="SF56F87CBJSXXXXXXXX"
                        />
                    </FlexColumn>
                    <Divider dashed style={{ marginBlock: 2 }} />
                    <Button onClick={onOk} type="primary">
                        Submit
                    </Button>
                </FlexColumn>
            </Card>
        </Modal>
    )
}
