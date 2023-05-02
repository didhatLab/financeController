import {GroupMember, GroupStackParams, SpendGroup} from "./types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {Spending} from "../../types/data-types";
import {BasicSpendList} from "../../components/user-spending";
import {StyleSheet, View} from "react-native";
import {GroupInfo} from "../../components/group-info";
import {FAB} from "react-native-paper";
import * as React from "react";


type Props = NativeStackScreenProps<GroupStackParams, 'GroupScreen'>


export function GroupSpendsScreen(props: Props) {
    const [groupSpends, setGroupSpends] = useState<Spending[]>([]);


    useEffect(() => {
        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/spending/group/get",
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ""
                    },
                    body: JSON.stringify({group_id: props.route.params.group.Id})

                }
            ))
            .then((response) => response.json())
            .then((json: Spending[]) => {
                setGroupSpends(json)
                console.log(json)
            })
            .catch((error) => {
                console.log(error)
            })

    }, [])

    return (
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'flex-start' }}>
            <GroupInfo group={props.route.params.group}
                       toParticipantList={(members: GroupMember[]) => props.navigation.navigate('MembersList', {members: members})}
                       toAdministration={(group: SpendGroup) => props.navigation.navigate('Admin', {group: props.route.params.group})}/>
            <BasicSpendList spends={groupSpends}
                            toNavigate={(spend: Spending) => props.navigation.navigate('EditGroupSpend', {spend: spend})}/>
            <FAB
                style={styles.fab}
                icon="plus"
                label="Add new note"
                onPress={() => props.navigation.navigate('AddNewGroupSpend', {
                    setGroupSpend: setGroupSpends,
                    groupId: props.route.params.group.Id
                })}

            />
        </View>
    )


}

const styles = StyleSheet.create({

    container: {

        flex: 1,

        backgroundColor: '#fff',

        paddingHorizontal: 10,

        paddingVertical: 20

    },

    titleContainer: {

        alignItems: 'center',

        justifyContent: 'center',

        flex: 1

    },

    title: {

        fontSize: 20

    },

    fab: {

        position: 'absolute',

        margin: 20,

        right: 0,

        bottom: 10

    }

});