import {RootStackParamList} from "../types/screen-types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";

type Props = NativeStackScreenProps<RootStackParamList, 'SplashScreen'>


export const SplashScreen = ({navigation}: Props) => {
    const [animating, setAnimating] = useState(true);


    useEffect(() => {
        console.log("kekke")
        AsyncStorage.getItem("token").then((value) => {

            if (value != null){
                navigation.navigate('DrawNavigationRoutes', {username: "dan"})
            } else {
                navigation.navigate("Login", {username: "kek"})
            }

        })
    })
    return <div>Loading...</div>
}
