import React from 'react';
// import styled from 'styled-components';
import styled from "styled-components";
import {Spending} from "../types/data-types";
import {ScrollView, StyleSheet, View, FlatList} from "react-native";
import {Avatar, Card, IconButton, Text, Button, List} from 'react-native-paper';
import Icon from '@mdi/react';
import {mdiAccount} from '@mdi/js';


const Header = styled.header`
  height: 50px;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  margin-top: 6px;
  padding: 0 11px;
`;

const Row = styled.text`
  align-items: center;
  flex-direction: row;
`;


const UserName = styled.text`
  padding-left: 8px;
  font-size: 14px;
  font-weight: bold;
  color: #010101;
`;

const PostDescription = styled.text`
  font-size: 14px;
  color: #222121;
  line-height: 16px;
  padding: 0 11px;
`;


export const Spend = (spend: Spending) => {
    return (
        <Card key={spend.Id} style={styles.container}>
            <Card.Title title={spend.Name}/>
            <Card.Content>
                <Text>{spend.Amount} {spend.Currency}</Text>
            </Card.Content>
            <Card.Actions>
                <Button>View</Button>
            </Card.Actions>
        </Card>
    );
};


export const SpendList = (props: { spends: Spending[] }) => {
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
