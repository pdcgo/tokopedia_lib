/* eslint-disable react-hooks/exhaustive-deps */
import {
    AnimatedProps,
    animated,
    useSpringRef,
    useTransition,
} from "@react-spring/web"
import { Breadcrumb, Button, Card, Input, Select } from "antd"
import React from "react"
import styled from "styled-components"
import { Flex, FlexColumn } from "../styled_components"

export type Props = {
    shopeeCatNames?: Array<string>
    productCount?: number
}

type Member = (
    props: AnimatedProps<{
        style: React.CSSProperties
    }>
) => React.ReactElement

const Box = styled.div`
    display: flex;
    flex: 1;
    position: relative;

    > div {
        position: absolute;
        display: flex;
        will-change: transform, opacity;
    }
`

export default function EtalaseMapCard(props: Props) {
    const inputs: Record<string, Member> = {
        input: ({ style }) => (
            <animated.div style={{ ...style, width: "100%" }}>
                <Input
                    onKeyUp={(e) => {
                        if (e.code == "Enter") {
                            setExpand("select")
                        }
                    }}
                    placeholder='Create new and hit "Enter"'
                />
            </animated.div>
        ),
        select: ({ style }) => (
            <animated.div style={{ ...style, width: "100%" }}>
                <Select
                    placeholder="Choose etalase"
                    style={{ width: "100%" }}
                />
            </animated.div>
        ),
    }

    const [expand, setExpand] = React.useState("select")

    const transRef = useSpringRef()
    const transitions = useTransition(expand, {
        ref: transRef,
        keys: null,
        from: {
            transform:
                expand == "select" ? "translateY(-50%)" : "translateY(50%)",
            opacity: 0,
        },
        enter: { transform: "translateY(0)", opacity: 1 },
        leave: {
            transform:
                expand == "select" ? "translateY(50%)" : "translateY(-50%)",
            opacity: 0,
        },
        initial: null,
        config: {
            duration: 200,
            velocity: 1
        }
    })

    React.useEffect(() => {
        transRef.start()
    }, [expand])

    return (
        <Card size="small" type="inner">
            <FlexColumn
                style={{ justifyContent: "space-between", height: "100%" }}
            >
                <Flex style={{ justifyContent: "space-between", flex: 1 }}>
                    <Breadcrumb
                        items={props.shopeeCatNames?.map((f) => ({ title: f }))}
                        separator=">"
                    />
                    <div style={{ width: "20px" }}></div>
                    <span style={{ flexShrink: 0 }}>
                        <strong>{props.productCount || 0}</strong> Product
                    </span>
                </Flex>
                <Flex>
                    <Box>
                        {transitions((style, ex) => {
                            const Element = inputs[ex]
                            return <Element style={style} />
                        })}
                    </Box>
                    <Button
                        onClick={() =>
                            setExpand((e) =>
                                e == "select" ? "input" : "select"
                            )
                        }
                        type={expand == "select" ? "primary" : "default"}
                    >
                        {expand == "select" ? "Or New" : "Cancel"}
                    </Button>
                </Flex>
            </FlexColumn>
        </Card>
    )
}
