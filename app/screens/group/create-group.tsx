import {NativeStackScreenProps} from "@react-navigation/native-stack";
import React, {useCallback, useContext, useState} from "react";
import {StyleSheet, View} from "react-native";
import {Box, CheckIcon, Input, Select, TextArea} from "native-base";
import {FAB} from "react-native-paper";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {GroupStackParams, SpendGroup} from "./types";
import {GroupContext} from "./context";

type Props = NativeStackScreenProps<GroupStackParams, 'CreateGroup'>


export function CreateNewGroup(props: Props) {
    const [groupTitle, setGroupTitle] = useState<string>('');
    const [groupDescription, setGroupDescription] = useState<string>('');

    const context = useContext(GroupContext)

    const onSaveNewGroup = () => {
        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/group/create",
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ""
                    },
                    body: JSON.stringify({group_name: groupTitle, group_description: groupDescription})
                }
            )).then((response) => response.json()).then((json) => {
            context.setGroups((groups) => {

                const newGroup: SpendGroup = {Id: json.group_id, Name: groupTitle, Description: groupDescription, Members: []}

                return [newGroup, ...groups]
            })

            props.navigation.goBack()
        })
    }


    return (
        <>
            <View style={styles.container}>
                <Box alignItems="center" justifyContent={"space-between"}>
                    <Input value={groupTitle} onChangeText={setGroupTitle} mx="3" placeholder="Title" w="100%"/>

                    <TextArea value={groupDescription}
                              onChangeText={text => setGroupDescription(text)} // for android and ios
                              w="100%" autoCompleteType={undefined} margin={10} placeholder='Description'/>
                </Box>
                <View style={styles.viewFab}>
                    <FAB style={styles.fab}
                         icon="cancel"
                         onPress={props.navigation.goBack}
                         size='small'
                    />
                    <FAB style={styles.fab} icon="check" size='small' onPress={onSaveNewGroup}/>
                </View>
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


