import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {HomeStackParams} from "./types";
import React, {useCallback, useState} from "react";
import {StyleSheet, View} from "react-native";
import {Box, CheckIcon, Input, Select, TextArea} from "native-base";
import {FAB} from "react-native-paper";
import AsyncStorage from "@react-native-async-storage/async-storage";

type Props = NativeStackScreenProps<HomeStackParams, 'ViewSpend'>


export function ViewSpending(props: Props) {
    const [spendTitle, setSpendTitle] = useState<string>(props.route.params.spend.Name);
    const [spendDescription, setSpendDescription] = useState<string>(props.route.params.spend.Description);
    const [amount, setAmount] = useState<string>(props.route.params.spend.Amount.toString());
    const [currency, setCurrency] = useState<string>(props.route.params.spend.Currency);

    const onSaveNote = useCallback(() => {
        console.log(props.route.params.spend.Id)
        AsyncStorage.getItem('token')
            .then((token) => {
                const dataForSend = {
                    spend_id: props.route.params.spend.Id,
                    name: spendTitle,
                    type: 'common',
                    amount: Number(amount),
                    description: spendDescription,
                }
                return fetch("http://localhost:4000/spending/update", {
                    method: 'POST',
                    body: JSON.stringify(dataForSend),
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ''
                    }
                })
            }).then((response) => response.json())


        props.navigation.goBack();
    }, [spendTitle, spendDescription, amount, currency])

    const onDeleteSpend = useCallback(() => {
        props.navigation.goBack()

        AsyncStorage.getItem('token')
            .then((token) => {
                const dataForSend = {
                    spend_id: props.route.params.spend.Id,
                }
                return fetch("http://localhost:4000/spending/delete", {
                    method: 'POST',
                    body: JSON.stringify(dataForSend),
                    headers: {
                        'Content-Type': 'application/json',
                        'Auth-Token': token ?? ''
                    }
                })
            }).then((_) => {
            let targetIndex = 0
            props.route.params.setSpends((spends) => {
                for (let i = 0; i < spends.length; i++){
                    if (spends[i].Id == props.route.params.spend.Id){
                        targetIndex = i;
                        break;
                    }
                }
                spends.splice(targetIndex, 1);
                return [...spends];
            })
        })

    }, [])


    return (
        <>
            <View style={styles.container}>
                <Box alignItems="center" justifyContent={"space-between"}>
                    <Input value={spendTitle} onChangeText={setSpendTitle} mx="3" placeholder="Title" w="100%"/>

                    <TextArea value={spendDescription}
                              onChangeText={text => setSpendDescription(text)} // for android and ios
                              w="100%" autoCompleteType={undefined} margin={10} placeholder='Description'/>

                    <Input value={amount} keyboardType="numeric" w='100%' onChangeText={setAmount}
                           placeholder="Amount"/>

                    <Select selectedValue={currency} minWidth="200" accessibilityLabel="Choose currency"
                            placeholder="Choose Currency" _selectedItem={{
                        bg: "teal.600",
                        endIcon: <CheckIcon size="5"/>
                    }} mt={1} onValueChange={itemValue => setCurrency(itemValue)} marginTop={10} isDisabled={true}>
                        <Select.Item label="RUB" value="RUB"/>
                        <Select.Item label="USD" value="USD"/>
                        <Select.Item label="EUR" value="EUR"/>
                    </Select>
                </Box>
                <View style={styles.viewFab}>
                    <FAB style={styles.fab}
                         icon="cancel"
                         onPress={props.navigation.goBack}
                         size='small'
                    />
                    <FAB style={styles.fab} icon='delete' size='small' onPress={() => onDeleteSpend()}/>
                    <FAB
                        style={styles.fab}
                        size='small'
                        icon="check"
                        onPress={() => onSaveNote()}
                    />
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


