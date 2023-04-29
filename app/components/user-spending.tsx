import React from 'react';
import {Spending} from "../types/data-types";
import {ScrollView, StyleSheet, View, FlatList} from "react-native";
import {Card, Text, Button, List} from 'react-native-paper';
import {HomeProps} from "../screens/menu/spend/types";


type OneSpendProps = {
    spend: Spending
    navigation: Pick<HomeProps, 'navigation'>
    setSpends: React.Dispatch<React.SetStateAction<Spending[]>>
}

export const Spend = (props: OneSpendProps) => {
    return (
        <Card key={props.spend.Id} style={styles.container}>
            <Card.Title title={props.spend.Name}/>
            <Card.Content>
                <Text>{props.spend.Amount} {props.spend.Currency}</Text>
            </Card.Content>
            <Card.Actions>
                <Button
                    onPress={() => props.navigation.navigation.navigate('ViewSpend', {
                        spend: props.spend,
                        setSpends: props.setSpends
                    })}>View</Button>
            </Card.Actions>
        </Card>
    );
};


type SpendsProps = {
    spends: Spending[]
    navigation: Pick<HomeProps, 'navigation'>
    setSpends: React.Dispatch<React.SetStateAction<Spending[]>>
}

export const SpendList = (props: SpendsProps) => {
    console.log(props.spends)
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

                    />)}
                // keyExtractor={item => '0' ? item.Id == undefined: item.Id.toString()}
            />
        </ScrollView>)
}


const styles = StyleSheet.create({
    container: {
        justifyContent: 'space-between',
        backgroundColor: 'grey',
        // padding: 20,
        margin: 10,
        textAlign: 'center',
        // marginTop: '0'
    },
    ff: {
        width: '100%'
    },
    listTitle: {
        fontSize: 20

    }
})
