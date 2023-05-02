import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {GroupStackParams, SpendGroup} from "./types";
import {useContext, useEffect} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {InflateUI} from "../../types/data-types";
import {GroupContext} from "./context";
import {SpendGroupList} from "../../components/spend-group";
import {FAB} from "react-native-paper";
import * as React from "react";
import {StyleSheet} from "react-native";

type Props = NativeStackScreenProps<GroupStackParams, 'GroupList'>

export function GroupsListScreen(props: Props) {

    const context = useContext(GroupContext)

    useEffect(() => {

        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/group/get",
                {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ""
                    }
                }
            ))
            .then((response) => response.json())
            .then((json: SpendGroup[]) => {
                context.setGroups(json)
                console.log(json)
            })
            .catch((error) => {
                console.log(error)
            })

    }, [])


    return (
        <>
            <SpendGroupList groups={context.groups} navigation={{navigation: props.navigation}}/>
            <FAB
                style={styles.fab}
                icon="plus"
                label="Add new group"
                onPress={() => props.navigation.navigate('CreateGroup', {})}

            />

        </>
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