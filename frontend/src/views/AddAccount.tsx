import React from 'react'
import { Card, Input, Alert, Button } from 'antd'
import { FlexColumn } from '../styled_components'

export default function AddAccount (): React.ReactElement {
  return (
    <FlexColumn>
      <Card  size='small' type='inner' title='Bulk Add Tokopedia Account'>
        <FlexColumn>
          <Alert
            type='info'
            message='Format: username|password|otp_secret'
            showIcon
          />
          <Input.TextArea
            size='large'
            autoSize={{ minRows: 24, maxRows: 24 }}
          />
          <Button type='primary'>Add Account</Button>
        </FlexColumn>
      </Card>
    </FlexColumn>
  )
}
