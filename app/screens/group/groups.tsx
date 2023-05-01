import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {GroupStackParams, SpendGroup} from "./types";
import {useContext, useEffect} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {InflateUI} from "../../types/data-types";
import {GroupContext} from "./context";
import {SpendGroupList} from "../../components/spend-group";

type Props = NativeStackScreenProps<GroupStackParams, 'GroupList'>

export function GroupsListScreen(props: Props){

    const context = useContext(GroupContext)

    useEffect(() => {

        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4000/group/get",
                {method: 'GET',
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
        <SpendGroupList groups={context.groups} navigation={{navigation: props.navigation}}/>
    )

}
