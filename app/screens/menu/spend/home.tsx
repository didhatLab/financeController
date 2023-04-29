import * as React from 'react';
import { Button, View, StyleSheet } from 'react-native';
import { createDrawerNavigator } from '@react-navigation/drawer';
import { NavigationContainer } from '@react-navigation/native';
import {DrawerParams} from "../types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import {InflateUI, Spending, CurrencyRate} from "../../../types/data-types";
import {CurrencyRateDisplay} from "../../../components/currency-rate";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {SpendList} from "../../../components/user-spending";
import {FAB} from "react-native-paper";
import {HomeProps} from "./types";


export function HomeScreen({ navigation }: HomeProps) {
    const [userSpends, setUserSpends] = useState<Spending[]>([]);
    const [currencyRate, setCurrencyRate] = useState<CurrencyRate>();


    useEffect(() => {

        AsyncStorage.getItem("token")
            .then((token) => fetch("http://localhost:4003/inflate_ui",
            {method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Auth-Token': token ?? ""
                }
            }
            ))
            .then((response) => response.json())
            .then((json: InflateUI) => {
                console.log(json.currency_rate.timestamp)
                setUserSpends(json.spends)
                setCurrencyRate(json.currency_rate)
            })
            .catch((error) => {
                console.log(error)
            })

    }, [])


    return (
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'flex-start' }}>
            <CurrencyRateDisplay base={currencyRate?.base} timestamp={currencyRate?.timestamp} rates={currencyRate?.rates}/>
            <SpendList spends={userSpends} navigation={{navigation: navigation}}/>
            <FAB
                style={styles.fab}
                icon="plus"
                label="Add new note"
                onPress={() => navigation.navigate('AddNewSpend', {})}

            />
        </View>
    );
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