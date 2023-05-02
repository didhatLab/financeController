import {GroupMember, SpendGroup} from "../screens/group/types";
import {Badge, Button, Card, Text} from "react-native-paper";
import React from "react";

type GroupInfoProps = {
    group: SpendGroup
    toParticipantList: (members: GroupMember[]) => any
    toAdministration: (group: SpendGroup) => any
}

export const GroupInfo = (props: GroupInfoProps) => {
    return ( <Card style={{width: '100%', alignItems: 'flex-start', justifyContent: 'flex-start'}}>
        <Card.Content>
            <Text variant='bodyMedium'>Group: {props.group.Name}</Text>
            <Text><Badge>{props.group.Members.length}</Badge> Members</Text>
            <Text>Description: {props.group.Description}</Text>
        </Card.Content>
        <Card.Actions>
            <Button onPress={() => props.toParticipantList(props.group.Members)}>Participant List</Button>
            <Button onPress={() => props.toAdministration(props.group)}>Group Administration</Button>
        </Card.Actions>
    </Card>
    )
}