import {FlatList, ScrollView, StyleSheet} from "react-native";
import {List} from "react-native-paper";
import {Divider} from "native-base";
import React from "react";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {GroupStackParams} from "./types";


type Props = NativeStackScreenProps<GroupStackParams, 'MembersList'>

export function MembersList(props: Props){

    return (
        <ScrollView showsVerticalScrollIndicator={false} style={styles.ff}>
            <FlatList
                data={props.route.params.members}
                renderItem={({item}) => (
                    <>
                        <List.Item
                            title={item.Username}
                            description={item.IsAdmin ? 'Admin': 'Member'}
                            descriptionNumberOfLines={1}
                            titleStyle={styles.listTitle}
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
