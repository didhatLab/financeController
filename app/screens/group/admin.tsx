import {FlatList, ScrollView, StyleSheet, View} from "react-native";
import {FAB, List} from "react-native-paper";
import {Box, CheckIcon, Divider, Input, Select, TextArea} from "native-base";
import React, {useContext, useEffect, useState} from "react";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import { Button } from 'react-native-paper';
import {GroupStackParams} from "./types";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {Spending} from "../../types/data-types";
import {GroupContext} from "./context";


type Props = NativeStackScreenProps<GroupStackParams, 'Admin'>


export function AdminGroup(props: Props){
    const [username, setUsername] = useState<string>('')

    const context = useContext(GroupContext)

    const addUser = () => {
            AsyncStorage.getItem("token")
                .then((token) => fetch("http://localhost:4000/group/member/add",
                    {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                            'Auth-Token': token ?? ""
                        },
                        body: JSON.stringify({group_id: props.route.params.group.Id, new_member_id: Number(username)})
                    }
                ))
    }

    const deleteUser = () => {
        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/group/member/delete",
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ""
                    },
                    body: JSON.stringify({group_id: props.route.params.group.Id, username_for_delete: username})
                }
            ))
    }

    const deleteGroup = () => {
        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/group/delete",
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ""
                    },
                    body: JSON.stringify({group_id: props.route.params.group.Id})
                }
            )).then((_) => {
                context.setGroups((groups) => {
                    let targetIndex = 0;
                    for (let i = 0; i < groups.length; ++i){
                        if (groups[i].Id == props.route.params.group.Id){
                            targetIndex = i
                            break
                        }
                    }
                    groups.splice(targetIndex, 1)

                    return [...groups]
                })

                props.navigation.navigate('GroupList', {})
        })
    }



    return (
        <>
            <View style={styles.container}>
                <Box alignItems="center" justifyContent={"space-between"}>
                    <Input value={username} onChangeText={setUsername} mx="3" placeholder="Title" w="100%"/>

                    <Button mode='contained' onPress={() => addUser()}>
                        Add user to Group
                    </Button>
                    <Button mode='contained' onPress={() => deleteUser()}>
                        Delete user from Group
                    </Button>

                    <Button mode='contained' onPress={() => deleteGroup()}>
                        Delete this Group
                    </Button>


                </Box>

            </View>

        </>

    );
}


const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        paddingHorizontal: 20,
        paddingVertical: 20

    },
    iconButton: {
        backgroundColor: 'rgba(46, 113, 102, 0.8)',
        position: 'absolute',
        right: 0,
        top: 40,
        margin: 10

    },

    title: {
        fontSize: 24,
        marginBottom: 20

    },
    text: {
        height: 300,
        fontSize: 16

    },

    fab: {
        marginRight: 20,
        marginBottom: 20,
        right: 0,
        bottom: 0,
        display: 'flex'
    },
    viewFab: {
        position: 'absolute',
        bottom: 0,
        right: 0,
        display: 'flex',
        justifyContent: 'center',
        flexDirection: 'row'

    }

});




