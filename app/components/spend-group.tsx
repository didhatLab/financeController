import {List} from "react-native-paper";
import {GroupPropsForRoute, SpendGroup} from "../screens/group/types";
import {GroupRouteProps} from "../screens/group/types";
import {FlatList, ScrollView, StyleSheet} from "react-native";
import {Divider} from "native-base";
import {useNavigation} from "@react-navigation/native";


type GroupProps = {
    group: SpendGroup
    navigation: Pick<GroupPropsForRoute, 'navigation'>
}

export const SpendGroupItem = (props: GroupProps) => {
    return (
        <>
            <List.Item title={props.group.Name} description={`${props.group.Members.length} participants`}
                       titleStyle={styles.listTitle}
                       onPress={() => props.navigation.navigation.navigate('GroupScreen', {group: props.group})}/>

            <Divider/>
        </>
    )
}

type GroupListProps = {
    groups: SpendGroup[]
    navigation: Pick<GroupPropsForRoute, 'navigation'>
}

export const SpendGroupList = (props: GroupListProps) => {

    return (
        <ScrollView showsVerticalScrollIndicator={false} style={styles.scroll}>
            <FlatList data={props.groups} renderItem={({item}) => (
                <SpendGroupItem group={item} navigation={props.navigation}/>
            )}/>
        </ScrollView>
    )
}




const styles = StyleSheet.create({
    listTitle: {
        fontSize: 20
    },
    scroll: {
        width: '100%'
    }
})
