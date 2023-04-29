import * as React from 'react';
import { Avatar, Card, IconButton, Text } from 'react-native-paper';
import {CurrencyRate} from "../types/data-types";
import {StyleSheet} from "react-native";

export const CurrencyRateDisplay = (rate: CurrencyRate) => (
    <Card.Content style={styles.container}>
        <Text variant="bodyMedium">{rate.base ?? "USD"}: {rate.rates ? rate.rates["RUB"]: 0} RUB: {rate.rates ? rate.rates["EUR"] : 0} EUR</Text>
    </Card.Content>
)

const styles = StyleSheet.create({
    container: {
        justifyContent: 'space-between',
        backgroundColor: 'grey',
        padding: 20,
        margin: 10,
        textAlign: 'center',
        width: '100%',
        marginTop: '0'
    }
})

