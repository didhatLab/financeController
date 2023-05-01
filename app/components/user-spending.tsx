import React from 'react';
import {Spending} from "../types/data-types";
import {ScrollView, StyleSheet, View, FlatList} from "react-native";
import {Card, Text, Button, List} from 'react-native-paper';
import {HomeProps} from "../screens/menu/spend/types";
import {useNavigation} from "@react-navigation/native";
import {Divider} from "native-base";


type SpendsProps = {
    spends: Spending[]
    navigation: Pick<HomeProps, 'navigation'>
    setSpends: React.Dispatch<React.SetStateAction<Spending[]>>
}

export const SpendList = (props: SpendsProps) => {
    return (
        <ScrollView showsVerticalScrollIndicator={false} style={styles.ff}>
            <FlatList
                data={props.spends}
                renderItem={({item}) => (
                    <List.Item
                        title={item.Name}
                        description={`${item.Amount} ${item.Currency}`}
                        descriptionNumberOfLines={1}
                        titleStyle={styles.listTitle}
                        onPress={() => props.navigation.navigation.navigate('ViewSpend', {
                            spend: item,
                            setSpends: props.setSpends
                        })}

                    />
                )}
            />
        </ScrollView>)
}

type BasicSpendsProps = {
    spends: Spending[]
    toNavigate: (spend: Spending) => void
}

export const BasicSpendList = (props: BasicSpendsProps) => {
    const navigation = useNavigation()

    return (
        <ScrollView showsVerticalScrollIndicator={false} style={styles.ff}>
            <FlatList
                data={props.spends}
                renderItem={({item}) => (
                    <>
                        <List.Item
                            title={item.Name}
                            description={`${item.Amount} ${item.Currency}`}
                            descriptionNumberOfLines={1}
                            titleStyle={styles.listTitle}
                            onPress={() => props.toNavigate(item)}

                        />
                        <Divider/>
                    </>
                )}
            />
        </ScrollView>)

}


const styles = StyleSheet.create({
    ff: {
        width: '100%'
    },
    listTitle: {
        fontSize: 20

    }
})
