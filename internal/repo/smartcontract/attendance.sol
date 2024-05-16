// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

contract AttendanceContract {
    struct AttendanceRecord {
        uint employeeId;
        uint index;
        uint checkInTime;
        uint checkOutTime;
        string details;
    }

    mapping(uint => AttendanceRecord[]) private attendanceRecords;
    mapping(address => uint256) private authorizedEntities;

    constructor() {
        authorizedEntities[msg.sender] = 1000000; // Deployer is authorized by default
    }

    // user ton tai
    modifier onlyAuthorized() {
        require(authorizedEntities[msg.sender] != 0, "Unauthorized access");
        _;
    }

    // truong hop la chinh chu user hoac admin thi dc tao
    function checkIn(uint _employeeId, uint _checkInTime, string memory _details) public onlyAuthorized {
        require( !(authorizedEntities[msg.sender] != 1000000  && _employeeId != authorizedEntities[msg.sender]), "Unauthorized access");

        uint _index = attendanceRecords[_employeeId].length;
        attendanceRecords[_employeeId].push(AttendanceRecord(_employeeId, _index, _checkInTime, 0, _details));
    }

    function checkOut(uint _employeeId, uint _index, uint _newCheckOutTime, string memory _newDetails) public onlyAuthorized {
        require(_index < attendanceRecords[_employeeId].length, "Index out of bounds");
        require( !(authorizedEntities[msg.sender] != 1000000  && _employeeId != authorizedEntities[msg.sender]), "Unauthorized access");

        attendanceRecords[_employeeId][_index].checkOutTime = _newCheckOutTime;
        attendanceRecords[_employeeId][_index].details =  string(abi.encodePacked(attendanceRecords[_employeeId][_index].details, ". checkout: ", _newDetails));
    }

    function getAttendanceByEmployeeId(uint _employeeId) public view returns (AttendanceRecord[] memory) {
        return attendanceRecords[_employeeId];
    }

    function getAttendanceByDateRange(uint _employeeId, uint _start, uint _end) public view returns (AttendanceRecord[] memory) {
        AttendanceRecord[] memory records = attendanceRecords[_employeeId];
        AttendanceRecord[] memory filteredRecords = new AttendanceRecord[](records.length);
        for (uint i = 0; i < records.length; i++) {
            if (records[i].checkInTime >= _start && records[i].checkInTime <= _end) {
                filteredRecords[i]=records[i];
            }
        }
        return filteredRecords;
    }

    function updateCheckIn(uint _employeeId, uint _index, uint _newCheckInTime, string memory _newDetails) public onlyAuthorized {
        require(_index < attendanceRecords[_employeeId].length, "Index out of bounds");
        if (authorizedEntities[msg.sender] != 1000000  && _employeeId != authorizedEntities[msg.sender]){
            return;
        }

        attendanceRecords[_employeeId][_index].checkInTime = _newCheckInTime;
        attendanceRecords[_employeeId][_index].details =  string(abi.encodePacked(attendanceRecords[_employeeId][_index].details, ". update checkin reason : ", _newDetails));
    }

    function authorizeEntity(address _entity, uint _employeeID) public onlyAuthorized {
        authorizedEntities[_entity] = _employeeID;
    }

    function revokeAuthorization(address _entity) public onlyAuthorized {
        authorizedEntities[_entity] = 0;
    }
}