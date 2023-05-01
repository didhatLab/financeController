import {GroupStackParams, SpendGroup} from "./types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {Spending} from "../../types/data-types";
import {BasicSpendList} from "../../components/user-spending";


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
        <BasicSpendList spends={groupSpends}
                        toNavigate={(spend: Spending) => props.navigation.navigate('EditGroupSpend', {spend: spend})}/>
    )


}
